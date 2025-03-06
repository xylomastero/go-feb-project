package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type ShutdownService struct {
	server *http.Server
	done   chan bool
}

func cancel(server *http.Server) *ShutdownService {
	return &ShutdownService{
		server: server,
		done:   make(chan bool),
	}
}

func (s *ShutdownService) HandleShutdown(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shutdown initiated"))

	// Signal shutdown
	go s.Shutdown()
}

func (s *ShutdownService) Shutdown() {
	log.Println("Shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	close(s.done)
}

func (s *ShutdownService) WaitForShutdown() {
	<-s.done
}
