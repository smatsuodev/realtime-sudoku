package main

import (
	"context"
	"log"
	"sudoku/config"
	"sudoku/infra/gorm"
	"sudoku/server"
)

func main() {
	cfg := config.NewEnvConfig()

	db, err := gorm.NewGormConnection()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.Run(ctx, cfg, db)
}
