package main

import (
	"log"
	"os"
	"strings"

	"def4alt/def4alt.com/internal/content"
	"def4alt/def4alt.com/internal/handler"
	"def4alt/def4alt.com/internal/render"
	"def4alt/def4alt.com/internal/server"
)

func main() {
	blog, err := content.Load("content")
	if err != nil {
		log.Fatalf("failed to load content: %v", err)
	}

	r, err := render.New(".")
	if err != nil {
		log.Fatalf("failed to load templates: %v", err)
	}

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8080"
	} else if !strings.Contains(addr, ":") {
		addr = ":" + addr
	}

	h := handler.New(r, blog)
	s := server.New(addr, h.Routes())

	log.Printf("Starting server on %s", addr)
	if err := s.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
