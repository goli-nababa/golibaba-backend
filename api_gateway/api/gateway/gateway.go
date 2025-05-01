package gateway

import (
	"api_gateway/api/http/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

// proxyRequest handles proxying an incoming HTTP request to the appropriate backend service.
// Parameters:
// - w: The HTTP response writer to send the response back to the client.
// - r: The incoming HTTP request from the client.
// - service: The backend service configuration where the request should be proxied.
func (g *Gateway) proxyRequest(w http.ResponseWriter, r *http.Request, service *types.RegisterRequest) error {
	// Extract the endpoint key by removing the service's base URL from the request path.
	endpointKey := strings.TrimPrefix(r.URL.Path, service.BaseUrl)
	log.Printf("Request path: %s, Endpoint key: %s", r.URL.Path, endpointKey)

	// Variables to store the matched endpoint and metadata.
	var matchedEndpoint types.Endpoint
	var exists bool
	var foundPattern string

	// Iterate over all endpoint mappings in the service configuration.
	// Check if the incoming request matches any endpoint pattern.
	for pattern, endpoint := range service.Mapping {

		// Use a helper function to match dynamic path patterns (e.g., "/users/{id}").
		if match := g.matchDynamicPath(endpointKey, pattern); match != nil {
			matchedEndpoint = endpoint
			exists = true
			foundPattern = pattern

			// Replace dynamic path variables with actual values in the target URL.
			endpointKey = g.replaceDynamicParts(matchedEndpoint.Url, pattern, endpointKey)
			break
		}
	}
	_ = foundPattern

	// If no matching endpoint is found, return a 404 error to the client.
	if !exists {
		http.Error(w, "Endpoint not found", http.StatusNotFound)
		return fmt.Errorf("endpoint %s not found in service %s", endpointKey, service.Name)
	}

	// Construct the target URL for the backend service using the service configuration.
	targetURL := fmt.Sprintf("http://%s:%s%s%s",
		service.Host,      // Backend service host.
		service.Port,      // Backend service port.
		service.UrlPrefix, // Optional prefix for the backend URL.
		endpointKey,       // Adjusted endpoint key.
	)

	log.Printf("Proxying to: %s", targetURL)

	// Create a new HTTP request to send to the backend service.
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return err
	}

	headers := make(map[string][]string)
	for key, values := range r.Header {
		headers[key] = values
	}

	modifiedHeaders, err := g.readUserInfo(headers)
	if err != nil {
		log.Printf("Error modifying headers: %v", err)
		http.Error(w, "Error processing headers", http.StatusInternalServerError)
		return err
	}

	// Add or override headers specific to the backend service configuration.
	for key, value := range service.Headers {
		req.Header.Set(key, value.(string))
	}

	// Set the modified headers on the proxied request
	for key, values := range modifiedHeaders {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Create an HTTP client with a timeout to send the request.
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error performing request: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Forward the backend service response status code to the client.
	w.WriteHeader(resp.StatusCode)

	// Stream the response body from the backend service to the client.
	_, err = io.Copy(w, resp.Body)
	return err
}

func (g *Gateway) readUserInfo(headers map[string][]string) (map[string][]string, error) {
	requestMeta := make(map[string]string)

	requestMeta["request_id"] = uuid.NewString()
	requestMeta["trace_id"] = uuid.NewString()

	if authHeaders, ok := headers["Authorization"]; ok && len(authHeaders) > 0 {

	}

	requestMetaString, err := json.Marshal(requestMeta)
	if err != nil {
		return nil, err
	}

	headers["X-User-Meta"] = []string{string(requestMetaString)}

	// Example: Remove a sensitive header
	delete(headers, "Authorization")

	return headers, nil
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
