package main

import (
	"context"
	"fmt"
	"log"
)

func run(ctx context.Context) error {
	port := 8080
	s, err := NewServer(port)
	if err != nil {
		return fmt.Errorf("failed run server: %v", err)
	}
	s.Run(ctx)
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("fatal application server: %v", err)
	}
}
