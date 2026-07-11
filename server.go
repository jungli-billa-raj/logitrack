package logitrack

import "net/http"

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
	s.router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logitrack server is up and running !!😇"))
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
