package usecase

import (
	"be/internal/repository"
	"fmt"
)

type JiraUsecase struct {
	repo *repository.JiraRepository
}

func NewJiraUsecase(repo *repository.JiraRepository) *JiraUsecase {
	return &JiraUsecase{repo: repo}
}

// Business logic for syncing Jira data
func (u *JiraUsecase) SyncJiraData() error {
	fmt.Println("SyncJiraData !!")

	data, err := u.repo.FetchJiraUpdates()
	if err != nil {
		return err
	}

	// Process Jira data (e.g., store in DB, send notifications)
	for _, item := range data {
		fmt.Println("Processing:", item.Title)
		err := u.repo.SaveJiraItem(item)
		if err != nil {
			fmt.Println("Failed to save:", item.Title)
		}
	}

	return nil
}
