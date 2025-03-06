// internal/usecase/sync_service.go
package usecase

import (
	"be/internal/repository"
	"log"
)

type JiraSync interface {
	ProcessSync() error
}

type jiraSync struct {
	syncRepo repository.SyncRepository
}

func NewJiraSyncUsecase(syncRepo repository.SyncRepository) JiraSync {
	return &jiraSync{syncRepo: syncRepo}
}

func (s *jiraSync) ProcessSync() error {
	// Fetch pending sync
	users, err := s.syncRepo.FetchUserList()
	if err != nil {
		return err
	}

	if err != nil || len(users) == 0 {
		log.Println("No user found. Err:", err)
		return nil
	}

	for _, user := range users {
		log.Println(user.JiraUserID)
	}

	return nil
}
