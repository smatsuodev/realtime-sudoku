package main

import (
	"context"
	"sudoku/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.Run(ctx)
}
