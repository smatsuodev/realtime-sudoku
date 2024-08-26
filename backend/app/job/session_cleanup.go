package job

import (
	"context"
	"log"
	"sudoku/repository"
	"time"
)

type SessionCleanupJob struct {
	sessionRepo repository.ISessionRepository
	interval    time.Duration
}

func NewSessionCleanupJob(sessionRepo repository.ISessionRepository, interval time.Duration) *SessionCleanupJob {
	return &SessionCleanupJob{
		sessionRepo: sessionRepo,
		interval:    interval,
	}
}

func (j *SessionCleanupJob) Run(ctx context.Context) {
	log.Println("Starting session cleanup job")

	ticker := time.NewTicker(j.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Session cleanup job stopped")
			return
		case <-ticker.C:
			if err := j.sessionRepo.DeleteExpired(); err != nil {
				log.Printf("Error during session cleanup: %v", err)
			} else {
				log.Println("Session cleanup completed successfully")
			}
		}
	}
}
