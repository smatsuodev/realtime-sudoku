package main

import (
	"context"
	"log"
	"sudoku/config"
	"sudoku/infra/gorm"
	"sudoku/job"
	"sudoku/server"
	"time"
)

// NOTE: ctx を正しく使えてない気がする
func main() {
	cfg := config.NewEnvConfig()

	db, err := gorm.NewGormConnection()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sessionCleanupJob := job.NewSessionCleanupJob(
		gorm.NewSessionRepository(db),
		time.Hour, // TODO: env から読む?
	)
	go sessionCleanupJob.Run(ctx)

	server.Run(ctx, cfg, db)
}
