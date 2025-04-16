package ucissue

import (
	"context"

	repository "be/internal/domain/repository"
)

func (uc *usecaseIssue) GetAssignedIssueByUserID(ctx context.Context, userID string) ([]repository.JiraIssueEntity, error) {

	user, err := uc.jiraDB.FetchUserByID(ctx, userID)
	if err != nil {
		uc.log.Println("FetchUserByID fail with err:", err)
		return nil, err
	}

	uc.log.Infof("GetAssignedIssueByUserID: userID: %d, email: %s", user.ID, user.Email)

	issues, err := uc.jiraDB.FetchJiraAssignedIssuesByEmail(ctx, user.Email)
	if err != nil {
		uc.log.Println("FetchUserList fail with err:", err)
		return nil, err
	}

	return issues, nil
}
