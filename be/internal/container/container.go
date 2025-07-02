package container

import (
	"be/internal/domain/config"
	repo "be/internal/domain/repository"
	"be/internal/domain/usecase"
	"be/internal/infrastructure/database"
	"be/internal/infrastructure/jira"
	jiradb "be/internal/repository/jira_db"
)

// Container holds all dependencies
type Container struct {
	Config         config.Config
	JiraRepository repo.JiraRepository
	JiraUseCase    usecase.JiraSyncUseCase
}

// NewContainer creates a new container with all dependencies
func NewContainer(cfg config.Config) (*Container, error) {
	// Initialize database connection
	db, err := database.NewConnection(cfg.GetDBConfig())
	if err != nil {
		return nil, err
	}

	// Initialize Jira client
	jiraClient, err := jira.NewClient(cfg.GetJiraConfig())
	if err != nil {
		return nil, err
	}

	// Initialize repository
	jiraRepo := jiradb.NewJiraRepository(db, jiraClient)

	// Initialize use case
	jiraUseCase := usecase.NewJiraSyncUseCase(jiraRepo)

	return &Container{
		Config:         cfg,
		JiraRepository: jiraRepo,
		JiraUseCase:    jiraUseCase,
	}, nil
}

// Close closes all connections and resources
func (c *Container) Close() error {
	// TODO: Implement cleanup logic
	return nil
}
