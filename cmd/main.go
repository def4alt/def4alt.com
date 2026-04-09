// Package main starts the blog server.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"def4alt/def4alt.com/internal/content"
	"def4alt/def4alt.com/internal/handler"
	"def4alt/def4alt.com/internal/render"
	"def4alt/def4alt.com/internal/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	blog, err := content.Load("content")
	if err != nil {
		return fmt.Errorf("load content: %w", err)
	}

	r, err := render.New(".")
	if err != nil {
		return fmt.Errorf("load templates: %w", err)
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
		return fmt.Errorf("server failed: %w", err)
	}
	return nil
}
