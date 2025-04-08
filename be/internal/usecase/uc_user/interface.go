package ucuser

import (
	repository "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	jiraDBRp "be/internal/repository/jira_db"
	"context"

	"github.com/sirupsen/logrus"
)

type UsecaseUser interface {
	GetAllJiraUsers(ctx context.Context) ([]repository.UserEntity, error)
}

type usecaseUser struct {
	cfg    *config.Config
	jiraDB jiraDBRp.JiraDBRepository
	log    *logrus.Logger
}

func NewUsecaseUser(cfg *config.Config, log *logrus.Logger, jiraDB jiraDBRp.JiraDBRepository) UsecaseUser {
	return &usecaseUser{
		cfg:    cfg,
		jiraDB: jiraDB,
		log:    log,
	}
}
