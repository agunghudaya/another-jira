package cron

import (
	"be/internal/usecase"
	"log"

	cron "github.com/robfig/cron/v3"
)

func RegisterJobs(c *cron.Cron, jiraUC *usecase.JiraUsecase) {

	_, err := c.AddFunc("@every 1m", func() { jiraUC.SyncJiraData() })
	if err != nil {
		log.Fatalf("Error registering cron job: %v", err)
	}

}

// New creates and returns a new cron scheduler
func New() *cron.Cron {
	return cron.New(cron.WithSeconds())
}
