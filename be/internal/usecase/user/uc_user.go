package ucuser

import (
	"context"

	repository "be/internal/domain/repository"
)

func (uc *usecaseUser) GetAllJiraUsers(ctx context.Context) ([]repository.UserEntity, error) {

	users, err := uc.jiraDB.FetchUserList(ctx)
	if err != nil {
		uc.log.Println("FetchUserList fail with err:", err)
		return nil, err
	}

	return users, nil
}
