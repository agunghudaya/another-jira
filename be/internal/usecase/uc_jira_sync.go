// internal/usecase/sync_service.go
package usecase

import (
	"be/internal/domain"
	"be/internal/infrastructure/config"
	"be/internal/repository"
	"database/sql"
	"log"
	"time"
)

type JiraSync interface {
	ProcessSync() error
	JiraUserSync(user *domain.User) error
}

type jiraSync struct {
	cfg      *config.Config
	db       *sql.DB
	syncRepo repository.SyncRepository
}

func NewJiraSyncUsecase(cfg *config.Config, db *sql.DB, syncRepo repository.SyncRepository) JiraSync {
	return &jiraSync{cfg: cfg, db: db, syncRepo: syncRepo}
}

func (s *jiraSync) ProcessSync() error {
	// Fetch pending sync
	users, err := s.syncRepo.FetchUserList()
	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return err
	}

	if len(users) == 0 {
		log.Println("No user found. Err:", err)
		return nil
	}

	for _, user := range users {
		s.JiraUserSync(&user)
	}

	return nil
}

func (s *jiraSync) JiraUserSync(user *domain.User) error {

	startedAt := time.Now()
	log.Printf("sync user_id\t:%s", user.JiraUserID)
	syncHistory, err := s.syncRepo.FetchPendingSync(user.JiraUserID)

	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return err
	}

	if syncHistory != nil {
		log.Println("Sync already in progress for user:", user.JiraUserID)
		return nil
	}

	jiraResponse, err := s.syncRepo.FetchJiraTasksWithFilter(user.JiraUserID, s.cfg)

	if err != nil {
		log.Println("FetchJiraTasksWithFilter fail with err:", err)
		s.syncRepo.InsertSyncHistory(s.db, user.JiraUserID, "fail", len(jiraResponse.Issues), jiraResponse.Total, err.Error(), startedAt)
		return err
	}

	s.syncRepo.InsertSyncHistory(s.db, user.JiraUserID, "success", len(jiraResponse.Issues), jiraResponse.Total, "", startedAt)

	return nil
}
