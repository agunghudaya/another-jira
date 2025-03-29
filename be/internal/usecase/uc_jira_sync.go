// internal/usecase/sync_service.go
package usecase

import (
	domainRP "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	"be/internal/repository"
	"context"
	"log"
	"time"

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

func (s *jiraSync) ProcessSync(ctx context.Context) error {

	users, err := s.syncRepo.FetchUserList(ctx)
	if err != nil {
		s.log.Println("FetchUserList fail with err:", err)
		return err
	}

	if len(users) == 0 {
		s.log.Println("No user found.")
		return nil
	}

	for _, user := range users {
		err := s.JiraUserSync(ctx, &user)
		if err != nil {
			s.log.Printf("Failed to sync user_id [%s]: %v", user.JiraUserID, err)
		}
	}

	return nil
}

func (s *jiraSync) JiraUserSync(ctx context.Context, user *domainRP.User) error {

	startedAt := time.Now()
	s.log.Printf("sync user_id\t:%s", user.JiraUserID)
	syncHistories, err := s.syncRepo.FetchPendingSync(ctx, user.JiraUserID)

	if err != nil {
		s.log.Errorln("FetchUserList fail with err:", err)
		return err
	}

	s.log.Printf("we have %d sync histories", len(syncHistories))

	totalExpectedRecords, records := 0, 0
	doSync := false

	if len(syncHistories) == 0 {
		doSync = true
	} else if len(syncHistories) > 0 {
		for _, sync := range syncHistories {
			s.log.Infof("\nlast sync\t:%s\ntotal\t:%d\ngot\t\t:%d",
				sync.CreatedAt.Format("02-01-2006 15:04:05"),
				sync.TotalExpectedRecords,
				sync.RecordsSynced)

			if sync.TotalExpectedRecords > totalExpectedRecords {
				totalExpectedRecords = sync.TotalExpectedRecords
			}

			records += sync.RecordsSynced
		}
	}

	if doSync {
		jiraResponse, err := s.syncRepo.FetchJiraTasksWithFilter(ctx, user.JiraUserID, s.cfg)

		if err != nil {
			log.Println("FetchJiraTasksWithFilter fail with err:", err)
			s.syncRepo.InsertSyncHistory(ctx, user.JiraUserID, "fail", len(jiraResponse.Issues), jiraResponse.Total, err.Error(), startedAt)
			return err
		}

		jiraIssues := domainRP.MapJiraResponseToJiraIssues(jiraResponse)
		s.log.Printf("user_id [%s] has %d issues", user.JiraUserID, len(jiraIssues))
		for _, issue := range jiraIssues {
			if err := s.syncRepo.InsertJiraIssues(ctx, issue); err != nil {
				s.log.Println("InsertJiraIssues fail with err:", err)
				return err
			}
		}

		s.syncRepo.InsertSyncHistory(ctx, user.JiraUserID, "success", len(jiraResponse.Issues), jiraResponse.Total, "", startedAt)

	} else if totalExpectedRecords == records {
		s.log.Printf("user_id [%s] is synced!", user.JiraUserID)
		return nil
	}

	return nil
}

func (s *jiraSync) GetJiraUserList(ctx context.Context) (user *[]domainRP.User, err error) {
	users, err := s.syncRepo.FetchUserList(ctx)
	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return nil, err
	}

	return &users, nil
}

func (s *jiraSync) CheckJiraSynced(ctx context.Context) error {
	users, err := s.syncRepo.FetchUserList(ctx)
	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return err
	}

	s.log.Printf("we have %d users", len(users))
	return nil
}
