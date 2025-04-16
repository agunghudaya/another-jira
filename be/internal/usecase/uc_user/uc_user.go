package ucuser

import (
	"context"

	repository "be/internal/domain/repository"
)

func (uc *usecaseUser) GetAllUsers(ctx context.Context) ([]repository.UserEntity, error) {

	users, err := uc.jiraDB.FetchUserList(ctx)
	if err != nil {
		uc.log.Println("FetchUserList fail with err:", err)
		return nil, err
	}

	return users, nil
}

func (uc *usecaseUser) GetUserByID(ctx context.Context, jiraID string) (repository.UserEntity, error) {

	user, err := uc.jiraDB.FetchUserByID(ctx, jiraID)
	if err != nil {
		uc.log.Println("FetchUserList fail with err:", err)
		return repository.UserEntity{}, err
	}

	return user, nil
}
