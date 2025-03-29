package repository

import (
	"be/internal/utils"
	"time"
)

type JiraIssue struct {
	ID                   int
	Key                  string
	Self                 string
	URL                  string
	AssigneeEmail        string
	AssigneeName         string
	ReporterEmail        string
	ReporterName         string
	CreatorEmail         string
	CreatorName          string
	Summary              string
	Description          string
	Created              time.Time
	Updated              time.Time
	DueDate              *time.Time
	StatusCategoryChange *time.Time
	TimeOriginalEstimate int
	IssueTypeName        string
	IssueTypeDescription string
	ProjectID            int
	ProjectKey           string
	ProjectName          string
	PriorityName         string
	TimeEstimate         int
	StatusName           string
	StatusDescription    string
	StatusCategoryName   string
	StatusCategoryKey    string
}

// MapJiraResponseToJiraIssues maps a Jira API response to a slice of JiraIssue structs.
func MapJiraResponseToJiraIssues(jiraResponse JiraResponse) []JiraIssue {
	var issues []JiraIssue

	for _, issue := range jiraResponse.Issues {

		created, err := utils.ParseStringToTime(issue.Fields.Created, time.RFC3339)
		if err != nil {
			created = time.Time{} // Handle parsing error (e.g., set to zero time)
		}

		updated, err := utils.ParseStringToTime(issue.Fields.Updated, time.RFC3339)
		if err != nil {
			updated = time.Time{} // Handle parsing error
		}

		issues = append(issues, JiraIssue{
			Key:  issue.Key,
			Self: issue.Self,
			//URL:                  issue.URL,
			AssigneeEmail: issue.Fields.Assignee.Email,
			AssigneeName:  issue.Fields.Assignee.DisplayName,
			ReporterEmail: issue.Fields.Reporter.Email,
			ReporterName:  issue.Fields.Reporter.DisplayName,
			CreatorEmail:  issue.Fields.Reporter.Email,
			CreatorName:   issue.Fields.Reporter.DisplayName,
			//Summary:              issue.Fields.Summary,
			Description: issue.Fields.Description,
			Created:     created,
			Updated:     updated,
			//DueDate:              issue.Fields.DueDate,
			//StatusCategoryChange: issue.Fields.StatusCategoryChangeDate,
			//TimeOriginalEstimate: issue.Fields.TimeOriginalEstimate,
			IssueTypeName:        issue.Fields.IssueType.Name,
			IssueTypeDescription: issue.Fields.IssueType.Description,
			//ProjectID:            issue.Fields.Project.ID,
			ProjectKey:  issue.Fields.Project.Key,
			ProjectName: issue.Fields.Project.Name,
			// PriorityName:         issue.Fields.Priority.Name,
			// TimeEstimate:         issue.Fields.TimeEstimate,
			// StatusName:           issue.Fields.Status.Name,
			// StatusDescription:    issue.Fields.Status.Description,
			// StatusCategoryName:   issue.Fields.Status.StatusCategory.Name,
			// StatusCategoryKey:    issue.Fields.Status.StatusCategory.Key,
		})
	}

	return issues
}
