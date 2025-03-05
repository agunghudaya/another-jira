// internal/usecase/sync_service.go
package usecase

import (
	"be/internal/repository"
	"log"
	"time"
)

type SyncService interface {
	ProcessSync() error
}

type syncService struct {
	syncRepo repository.SyncRepository
}

func NewSyncService(syncRepo repository.SyncRepository) SyncService {
	return &syncService{syncRepo: syncRepo}
}

func (s *syncService) ProcessSync() error {
	// Fetch pending sync
	sync, err := s.syncRepo.FetchPendingSync()
	if err != nil {
		return err
	}

	if sync == nil {
		log.Println("No pending sync found.")
		return nil
	}

	log.Printf("Processing sync for Jira ID: %s\n", sync.JiraID)

	// Simulate data fetching
	time.Sleep(3 * time.Second) // Replace with actual Jira API calls

	// Simulate success
	err = s.syncRepo.MarkSyncAsCompleted(sync.ID, true, 50, nil)
	if err != nil {
		log.Fatalf("Failed to update sync status: %v", err)
	}

	log.Println("Sync completed successfully.")
	return nil
}
