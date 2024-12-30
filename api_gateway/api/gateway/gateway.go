package gateway

import (
	"api_gateway/api/http/types"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/goli-nababa/golibaba-backend/modules/cache"
)

// Gateway represents the API gateway.
type Gateway struct {
	cacheProvider *cache.ObjectCache[*types.RegisterRequest]
}

// NewGateway creates a new Gateway instance.
func NewGateway(provider *cache.ObjectCache[*types.RegisterRequest]) *Gateway {
	return &Gateway{
		cacheProvider: provider,
	}
}

// ServeHTTP implements the http.Handler interface.
func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Match the request path to a service
	service, err := g.matchService(r.Context(), r.URL.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Proxy the request
	if err := g.proxyRequest(w, r, service); err != nil {
		http.Error(w, "Failed to process the request", http.StatusBadGateway)
		log.Printf("Error proxying request: %v", err)
	}
}

// matchService finds the appropriate service based on the request path.
func (g *Gateway) matchService(ctx context.Context, path string) (*types.RegisterRequest, error) {
	// Extract the service name and version from the path
	segments := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(segments) < 1 {
		return nil, errors.New("invalid request path")
	}

	// Construct the expected key pattern: service.{service_name}.{service_version}
	serviceVersion := segments[0]
	serviceName := segments[1]
	cacheKey := fmt.Sprintf("%s.%s", serviceName, serviceVersion)

	// Check if the service exists
	exists, err := g.cacheProvider.Exists(ctx, cacheKey)

	if err != nil {
		log.Printf("Error checking service existence: %v", err)
		return nil, errors.New("failed to check service existence")
	}
	if !exists {
		return nil, errors.New("service not found")
	}

	// Fetch the service details
	service, err := g.cacheProvider.Get(ctx, cacheKey)
	if err != nil {
		log.Printf("Error fetching service %s: %v", cacheKey, err)
		return nil, errors.New("failed to fetch service data")
	}

	/*	// Verify if the path matches the service's URL prefix
		if !strings.HasPrefix(path, service.UrlPrefix) {
			return nil, errors.New("path does not match service URL prefix")
		}*/

	return service, nil
}

func (g *Gateway) matchDynamicPath(actualPath, pattern string) map[string]string {
	actual := strings.Split(strings.Trim(actualPath, "/"), "/")
	expected := strings.Split(strings.Trim(pattern, "/"), "/")

	if len(actual) != len(expected) {
		return nil
	}

	params := make(map[string]string)
	for i, exp := range expected {
		if strings.HasPrefix(exp, "{") && strings.HasSuffix(exp, "}") {
			paramName := strings.Trim(exp, "{}")
			params[paramName] = actual[i]
		} else if exp != actual[i] {
			return nil
		}
	}
	return params
}

func (g *Gateway) replaceDynamicParts(targetPattern, sourcePattern, actualPath string) string {
	sourceParts := strings.Split(strings.Trim(sourcePattern, "/"), "/")
	actualParts := strings.Split(strings.Trim(actualPath, "/"), "/")
	targetParts := strings.Split(strings.Trim(targetPattern, "/"), "/")

	result := make([]string, len(targetParts))
	paramMap := make(map[string]string)

	// Build parameter mapping
	for i, part := range sourceParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			paramName := strings.Trim(part, "{}")
			paramMap[paramName] = actualParts[i]
		}
	}

	// Replace parameters in target
	for i, part := range targetParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			paramName := strings.Trim(part, "{}")
			if value, exists := paramMap[paramName]; exists {
				result[i] = value
			} else {
				result[i] = part // Keep original if no mapping found
			}
		} else {
			result[i] = part
		}
	}

	return "/" + strings.Join(result, "/")
}

func (g *Gateway) proxyRequest(w http.ResponseWriter, r *http.Request, service *types.RegisterRequest) error {
	endpointKey := strings.TrimPrefix(r.URL.Path, service.BaseUrl)
	log.Printf("Request path: %s, Endpoint key: %s", r.URL.Path, endpointKey)

	var matchedEndpoint types.Endpoint
	var exists bool
	var foundPattern string

	for pattern, endpoint := range service.Mapping {
		if match := g.matchDynamicPath(endpointKey, pattern); match != nil {
			matchedEndpoint = endpoint
			exists = true
			foundPattern = pattern
			// Replace dynamic parts in the target URL
			endpointKey = g.replaceDynamicParts(matchedEndpoint.Url, pattern, endpointKey)
			break
		}
	}
	_ = foundPattern

	if !exists {
		http.Error(w, "Endpoint not found", http.StatusNotFound)
		return fmt.Errorf("endpoint %s not found in service %s", endpointKey, service.Name)
	}

	targetURL := fmt.Sprintf("http://%s:%s%s%s",
		service.Host,
		service.Port,
		service.UrlPrefix,
		endpointKey)

	log.Printf("Proxying to: %s", targetURL)
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return err
	}

	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	for key, value := range service.Headers {
		req.Header.Set(key, value.(string)) // Assuming the value is a string
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error performing request: %v", err)
		return err
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	return err
}
func (g *Gateway) validatePermission(permissionList map[string]any, userRole string, serviceName string, endpointKey string) bool {
	if userRole == "" {
		log.Printf("Missing user role in request")
		return false
	}

	permissions, exists := permissionList[userRole]
	if !exists {
		log.Printf("No permissions defined for role %s", userRole)
		return false
	}

	permissionSlice, ok := permissions.([]string)
	if !ok {
		log.Printf("Invalid permission format for role %s", userRole)
		return false
	}

	requiredPermission := fmt.Sprintf("%s:%s", serviceName, endpointKey)
	for _, permission := range permissionSlice {
		if permission == requiredPermission {
			return true
		}
	}

	log.Printf("Permission %s not found for role %s", requiredPermission, userRole)
	return false
}
