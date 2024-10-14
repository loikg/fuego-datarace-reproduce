package main

import (
	"fmt"
	"testing"

	"github.com/go-fuego/fuego"
)

func TestDataRace(t *testing.T) {
	t.Parallel()

	for n := range 2 {
		t.Run(fmt.Sprintf("create server %d", n), func(t *testing.T) {
			t.Parallel()
			createFuegoServer(t)
		})
	}
}

func createFuegoServer(t *testing.T) {
	t.Helper()

	server := fuego.NewServer()
	fuego.Get(server, "/", func(ctx *fuego.ContextNoBody) (any, error) {
		return nil, nil
	})
}
