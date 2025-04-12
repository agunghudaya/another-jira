package jirasync

import (
	"be/internal/infrastructure/config"
	"be/internal/infrastructure/logger"
	"context"

	repository "be/internal/domain/repository"
	jiraAtlassianRp "be/internal/repository/jira_atlassian"
	jiraDBRp "be/internal/repository/jira_db"
)

type JiraSync interface {
	CheckJiraSynced(ctx context.Context) error
	GetJiraUserList(ctx context.Context) (user *[]repository.UserEntity, err error)
	JiraUserSync(ctx context.Context, user *repository.UserEntity) error
	ProcessSync(ctx context.Context) error
}

type jiraSync struct {
	cfg           *config.Config
	jiraAtlassian jiraAtlassianRp.JiraAtlassianRepository
	jiraDB        jiraDBRp.JiraDBRepository
	log           logger.Logger
}

func NewJiraSyncUsecase(cfg *config.Config, log logger.Logger, jiraDB jiraDBRp.JiraDBRepository, jiraAtlassian jiraAtlassianRp.JiraAtlassianRepository) JiraSync {
	return &jiraSync{
		cfg:           cfg,
		jiraAtlassian: jiraAtlassian,
		jiraDB:        jiraDB,
		log:           log,
	}
}
