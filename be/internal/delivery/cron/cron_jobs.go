package cron

import (
	"be/internal/usecase"
	"context"

	cron "github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	cron       *cron.Cron
	jiraSyncUC usecase.JiraSync
	logger     *logrus.Logger
}

// NewWorker initializes a new cron job worker
func NewWorker(logger *logrus.Logger, jiraSyncUC usecase.JiraSync) *Worker {
	return &Worker{
		logger:     logger,
		jiraSyncUC: jiraSyncUC,
		cron:       cron.New(),
	}
}

// Start registers and starts cron jobs
func (w *Worker) Start(ctx context.Context) {

	_, err := w.cron.AddFunc("* */2 * * *", func() { w.jiraSyncUC.CheckJiraSynced(ctx) })
	if err != nil {
		w.logger.Println("Error scheduling cron job:", err)
		return
	}

	w.cron.Start()
	w.logger.Println("Cron jobs started.")
}
