package ucuser

import (
	repository "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	"be/internal/infrastructure/logger"

	jiraDBRp "be/internal/repository/jira_db"
	"context"
)

type UsecaseUser interface {
	GetAllUsers(ctx context.Context) ([]repository.UserEntity, error)
	GetUserByID(ctx context.Context, jiraID string) (repository.UserEntity, error)
}

type usecaseUser struct {
	cfg    config.Config
	jiraDB jiraDBRp.JiraDBRepository
	log    logger.Logger
}

func NewUsecaseUser(cfg config.Config, log logger.Logger, jiraDB jiraDBRp.JiraDBRepository) UsecaseUser {
	return &usecaseUser{
		cfg:    cfg,
		jiraDB: jiraDB,
		log:    log,
	}
}
