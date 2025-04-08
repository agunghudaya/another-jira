package ucjirasync

import (
	repository "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	jiraAtlassianRp "be/internal/repository/jira_atlassian"
	jiraDBRp "be/internal/repository/jira_db"
	"context"

	"github.com/sirupsen/logrus"
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
	log           *logrus.Logger
}

func NewJiraSyncUsecase(cfg *config.Config, log *logrus.Logger, jiraDB jiraDBRp.JiraDBRepository, jiraAtlassian jiraAtlassianRp.JiraAtlassianRepository) JiraSync {
	return &jiraSync{
		cfg:           cfg,
		jiraAtlassian: jiraAtlassian,
		jiraDB:        jiraDB,
		log:           log,
	}
}
