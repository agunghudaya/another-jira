package ucjirasync

import (
	"context"
	"log"
	"time"

	repository "be/internal/domain/repository"
)

func (s *jiraSync) ProcessSync(ctx context.Context) error {

	users, err := s.jiraDB.FetchUserList(ctx)
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

func (s *jiraSync) JiraUserSync(ctx context.Context, user *repository.UserEntity) error {
	startedAt := time.Now()
	s.log.Printf("sync user_id\t:%s", user.JiraUserID)

	syncHistories, err := s.fetchSyncHistories(ctx, user.JiraUserID)
	if err != nil {
		return err
	}

	doSync, totalExpectedRecords, records := s.analyzeSyncHistories(syncHistories)

	if doSync {
		return s.performSync(ctx, user, startedAt)
	} else if totalExpectedRecords == records {
		s.log.Printf("user_id [%s] is synced!", user.JiraUserID)
		return nil
	}

	return nil
}

func (s *jiraSync) fetchSyncHistories(ctx context.Context, jiraUserID string) ([]repository.SyncHistory, error) {
	syncHistories, err := s.jiraDB.FetchPendingSync(ctx, jiraUserID)
	if err != nil {
		s.log.Errorln("FetchPendingSync fail with err:", err)
		return nil, err
	}
	return syncHistories, nil
}

func (s *jiraSync) analyzeSyncHistories(syncHistories []repository.SyncHistory) (bool, int, int) {
	totalExpectedRecords, records := 0, 0
	doSync := len(syncHistories) == 0

	for _, sync := range syncHistories {
		if sync.TotalExpectedRecords > totalExpectedRecords {
			totalExpectedRecords = sync.TotalExpectedRecords
		}
		records += sync.RecordsSynced
	}

	return doSync, totalExpectedRecords, records
}

func (s *jiraSync) performSync(ctx context.Context, user *repository.UserEntity, startedAt time.Time) error {
	jiraResponse, err := s.jiraAtlassian.FetchJiraTasksWithFilter(ctx, user.JiraUserID, s.cfg)
	if err != nil {
		s.log.Println("FetchJiraTasksWithFilter fail with err:", err)
		return err
	}

	jiraIssues := repository.MapJiraResponseToJiraIssues(jiraResponse)
	s.log.Printf("user_id [%s] has %d issues", user.JiraUserID, len(jiraIssues))

	for _, issue := range jiraIssues {
		if err := s.processJiraIssue(ctx, issue); err != nil {
			return err
		}
	}

	return s.jiraDB.InsertSyncHistory(ctx, user.JiraUserID, "success", len(jiraResponse.Issues), jiraResponse.Total, "", startedAt)
}

func (s *jiraSync) processJiraIssue(ctx context.Context, issue repository.JiraIssueEntity) error {
	existingIssue, err := s.jiraDB.FetchJiraIssue(ctx, issue.Key)
	if err != nil {
		s.log.Println("FetchJiraIssue fail with err:", err)
		return err
	}

	if existingIssue.Key == issue.Key {
		if issue.Updated.After(existingIssue.Updated) {
			if err := s.jiraDB.UpdateJiraIssue(ctx, issue); err != nil {
				s.log.Infoln("UpdateJiraIssue fail with err:", err)
				return err
			}

		} else {
			s.log.Infof("Skipping issue %s as it already exists and got no update since %s", issue.Key, issue.Updated.Format("02-01-2006 15:04:05"))
			return nil
		}
	}

	if err := s.jiraDB.InsertJiraIssue(ctx, issue); err != nil {
		s.log.Infoln("InsertJiraIssues fail with err:", err)
		return err
	}

	if err := s.updateJiraIssueHistories(ctx, issue); err != nil {
		s.log.Infoln("updateJiraIssueHistories fail with err:", err)
		return err
	}

	return nil
}

func (s *jiraSync) updateJiraIssueHistories(ctx context.Context, issue repository.JiraIssueEntity) error {

	histories, err := s.jiraAtlassian.FetchJiraIssueHistories(ctx, issue.Key, s.cfg)
	if err != nil {
		s.log.Println("FetchJiraIssueHistories fail with err:", err)
		return err
	}

	historyEntities := repository.MapToJiraIssueHistoryEntities(histories)

	for _, historyEntity := range historyEntities {
		err := s.jiraDB.InsertJiraIssueHistory(ctx, historyEntity)
		if err != nil {
			log.Println("InsertJiraIssueHistory fail with err:", err)
			return err
		}
	}

	return nil
}

func (s *jiraSync) GetJiraUserList(ctx context.Context) (user *[]repository.UserEntity, err error) {
	users, err := s.jiraDB.FetchUserList(ctx)
	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return nil, err
	}

	return &users, nil
}

func (s *jiraSync) CheckJiraSynced(ctx context.Context) error {
	users, err := s.jiraDB.FetchUserList(ctx)
	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return err
	}

	s.log.Printf("we have %d users", len(users))
	return nil
}
