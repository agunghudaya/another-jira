package repository

import (
	"be/internal/utils"
	"time"
)

type JiraIssueHistoryEntity struct {
	ID       int
	IssueKey string
	Field    string
	Oldvalue string
	NewValue string

	Created time.Time
}

func MapToJiraIssueHistoryEntities(jiraResponse JiraIssueHistoryResponse) []JiraIssueHistoryEntity {
	var histories []JiraIssueHistoryEntity
	format := "2006-01-02T15:04:05.000-0700"

	for _, history := range jiraResponse.Changelog.Histories {

		created, err := utils.ParseStringToTime(history.Created, format)
		if err != nil {
			created = time.Time{} // Handle parsing error (e.g., set to zero time)
		}

		historyEntity := JiraIssueHistoryEntity{
			Created:  created,
			IssueKey: jiraResponse.Key,
		}

		for _, items := range history.Items {
			historyEntity.Field = items.Field
			historyEntity.Oldvalue = items.FromString
			historyEntity.NewValue = items.ToString
		}

		histories = append(histories, historyEntity)
	}

	return histories
}
