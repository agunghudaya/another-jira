// internal/usecase/sync_service.go
package jira_sync

import (
	domainRP "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	jiraAtlassianRp "be/internal/repository/jira_atlassian"
	jiraDBRp "be/internal/repository/jira_db"
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
	cfg           *config.Config
	log           *logrus.Logger
	jiraDB        jiraDBRp.JiraDBRepository
	jiraAtlassian jiraAtlassianRp.JiraAtlassianRepository
}

func NewJiraSyncUsecase(cfg *config.Config, log *logrus.Logger, jiraDB jiraDBRp.JiraDBRepository, jiraAtlassian jiraAtlassianRp.JiraAtlassianRepository) JiraSync {
	return &jiraSync{
		cfg:           cfg,
		log:           log,
		jiraDB:        jiraDB,
		jiraAtlassian: jiraAtlassian,
	}
}
