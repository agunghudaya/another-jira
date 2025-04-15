package cron

import (
	"be/internal/infrastructure/logger"
	jirasync "be/internal/usecase/uc_jira_sync"
	"context"

	cron "github.com/robfig/cron/v3"
)

type Worker struct {
	cron       *cron.Cron
	jiraSyncUC jirasync.JiraSync
	log        logger.Logger
}

// NewWorker initializes a new cron job worker
func NewWorker(log logger.Logger, jiraSyncUC jirasync.JiraSync) *Worker {
	return &Worker{
		log:        log,
		jiraSyncUC: jiraSyncUC,
		cron:       cron.New(),
	}
}

// Start registers and starts cron jobs
func (w *Worker) Start(ctx context.Context) {

	_, err := w.cron.AddFunc("*/2 * * * *", func() { w.jiraSyncUC.ProcessSync(ctx) })
	if err != nil {
		w.log.Println("Error scheduling cron job:", err)
		return
	}

	w.cron.Start()
	w.log.Println("Cron jobs started.")
}
