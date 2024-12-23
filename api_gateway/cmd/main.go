package main

import (
	"api_gateway/app"
	"api_gateway/config"
	"flag"
	"os"

	http "api_gateway/api/http"
)

/*
// Service represents a registered microservice.
type Service struct {
	Name           string            `json:"name"`
	URLMap         map[string]string `json:"url_map"`
	Rules          map[string]string `json:"rules"`
	HeartbeatURL   string            `json:"heartbeat_url"`
	PermissionJSON string            `json:"permissions"`
}

type Gateway struct {
	Services       map[string]*Service
	ServiceState   map[string]bool
	Mutex          sync.RWMutex
	RedisClient    *redis.Client
	UserServiceURL string
}

func NewGateway(redisAddr, userServiceURL string) *Gateway {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &Gateway{
		Services:       make(map[string]*Service),
		ServiceState:   make(map[string]bool),
		RedisClient:    redisClient,
		UserServiceURL: userServiceURL,
	}
}

func (g *Gateway) RegisterService(w http.ResponseWriter, r *http.Request) {
	var service Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	g.Mutex.Lock()
	g.Services[service.Name] = &service
	g.ServiceState[service.Name] = true
	g.Mutex.Unlock()

	w.WriteHeader(http.StatusOK)
}

func (g *Gateway) MonitorHeartbeat() {
	for {
		g.Mutex.RLock()
		for name, service := range g.Services {
			go func(name string, service *Service) {
				resp, err := http.Get(service.HeartbeatURL)
				g.Mutex.Lock()
				defer g.Mutex.Unlock()
				if err != nil || resp.StatusCode != http.StatusOK {
					g.ServiceState[name] = false
					log.Printf("Service %s is unavailable", name)
				} else {
					g.ServiceState[name] = true
				}
			}(name, service)
		}
		g.Mutex.RUnlock()
		time.Sleep(10 * time.Second)
	}
}

func (g *Gateway) CheckAvailability(name string) bool {
	g.Mutex.RLock()
	defer g.Mutex.RUnlock()
	return g.ServiceState[name]
}

func (g *Gateway) ProxyRequest(w http.ResponseWriter, r *http.Request) {
	serviceName := mux.Vars(r)["service"]
	url := mux.Vars(r)["url"]

	if !g.CheckAvailability(serviceName) {
		go g.UndemandCheck(serviceName)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	service, ok := g.Services[serviceName]
	if !ok {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	proxyURL, ok := service.URLMap[url]
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, proxyURL, http.StatusTemporaryRedirect)
}

func (g *Gateway) UndemandCheck(name string) {
	service, ok := g.Services[name]
	if !ok {
		return
	}

	resp, err := http.Get(service.HeartbeatURL)
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	if err == nil && resp.StatusCode == http.StatusOK {
		g.ServiceState[name] = true
		log.Printf("Service %s is now available", name)
	} else {
		g.ServiceState[name] = false
		log.Printf("Service %s is still unavailable", name)
	}
}*/

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	/*gateway := NewGateway("localhost:6379", "http://user-service")
	defer gateway.RedisClient.Close()

	router := mux.NewRouter()
	// api-gateway:8080/register
	router.HandleFunc("/register", gateway.RegisterService).Methods("POST")
	router.HandleFunc("/proxy/{service}/{url}", gateway.ProxyRequest).Methods("GET")

	go gateway.MonitorHeartbeat()

	log.Println("API Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))*/
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.MustNewApp(c)

	err := http.Bootstrap(appContainer, c.Server)

	if err != nil {
		return
	}
}
