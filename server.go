package logitrack

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Server struct {
	repo   LogisticsRepository
	router *http.ServeMux
}

func NewServer(repo LogisticsRepository) *Server {
	mux := http.NewServeMux()
	s := &Server{
		repo:   repo,
		router: mux,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("GET /health", s.handleHealth())
	s.router.HandleFunc("POST /shipments", s.handleDispatchShipment())
	s.router.HandleFunc("GET /warehouses/{id}/inventory", s.handleGetInventory())
	s.router.HandleFunc("GET /shipments/{id}", s.handleGetShipment())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 1. Pass the request along to the next person in line (our router)
		next.ServeHTTP(w, r)

		// 2. Once the handler finishes, log the details to the console!
		log.Printf("Method: %s | Path: %s | Duration: %s", r.Method, r.URL.Path, time.Since(start))
	})
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loggingMiddleware(s.router).ServeHTTP(w, r)
}

// endpoints

func (s *Server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logitrack server is up and running !!😇"))
	}
}

func (s *Server) handleDispatchShipment() http.HandlerFunc {
	type response struct {
		ShipmentID string `json:"shipment_id,omitempty"`
		Error      string `json:"error,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		//Set Headers
		w.Header().Set("Content-Type", "application/json")

		var order DispatchOrder
		//Decode the incoming request
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response{Error: "Invalid JSON request"})
			return
		}

		//Business Logic
		shipmentID, err := s.repo.DispatchShipment(r.Context(), order)
		if err != nil {
			// we can check SQL state here. But rn we are simply returning 400 Bad Request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response{Error: err.Error()})
			return
		}
		//Return success state with the newID
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response{ShipmentID: shipmentID})
	}
}

func (s *Server) handleGetInventory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Grab the {id} wildcard value from the incoming URL Path
		warehouseId := r.PathValue("id")
		if warehouseId == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing Warehouse ID"})
			return
		}

		// 2. Query DB
		items, err := s.repo.GetWarehouseInventory(r.Context(), warehouseId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)

	}
}

func (s *Server) handleGetShipment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. extract the shipmentID from URL Path
		shipmentID := r.PathValue("id")
		if shipmentID == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "ShipmentID not provided"})
			return
		}

		// 2. Query the DB
		status, err := s.repo.GetShipment(r.Context(), shipmentID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(status)
	}
}
