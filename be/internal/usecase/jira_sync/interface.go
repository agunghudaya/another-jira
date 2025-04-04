// internal/usecase/sync_service.go
package jira_sync

import (
	domainRP "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	"be/internal/repository"
	"context"

	"github.com/sirupsen/logrus"
)

type JiraSync interface {
	CheckJiraSynced(ctx context.Context) error
	GetJiraUserList(ctx context.Context) (user *[]domainRP.User, err error)
	JiraUserSync(ctx context.Context, user *domainRP.User) error
	ProcessSync(ctx context.Context) error
}

type jiraSync struct {
	cfg      *config.Config
	log      *logrus.Logger
	syncRepo repository.SyncRepository
}

func NewJiraSyncUsecase(cfg *config.Config, log *logrus.Logger, syncRepo repository.SyncRepository) JiraSync {
	return &jiraSync{
		cfg:      cfg,
		log:      log,
		syncRepo: syncRepo,
	}
}
