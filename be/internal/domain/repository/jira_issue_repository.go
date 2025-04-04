package repository

import (
	"be/internal/utils"
	"log"
	"time"
)

type JiraIssue struct {
	ID                   int
	Key                  string
	AssigneeEmail        string
	AssigneeName         string
	Created              time.Time
	CreatorEmail         string
	CreatorName          string
	Description          string
	DueDate              *time.Time
	IssueTypeDescription string
	IssueTypeName        string
	PriorityName         string
	ProjectID            int
	ProjectKey           string
	ProjectName          string
	ReporterEmail        string
	ReporterName         string
	Self                 string
	StatusCategoryChange *time.Time
	StatusCategoryKey    string
	StatusCategoryName   string
	StatusDescription    string
	StatusName           string
	Summary              string
	Updated              time.Time
	URL                  string

	//time estimation
	AggregateTimeEstimate         float64
	AggregateTimeOriginalEstimate float64
	TimeEstimate                  float64
	TimeOriginalEstimate          float64
}

// MapJiraResponseToJiraIssues maps a Jira API response to a slice of JiraIssue structs.
func MapJiraResponseToJiraIssues(jiraResponse JiraIssueResponse) []JiraIssue {
	var issues []JiraIssue

	format := "2006-01-02T15:04:05.000-0700"

	for _, issue := range jiraResponse.Issues {

		created, err := utils.ParseStringToTime(issue.IssueFields.Created, format)
		if err != nil {
			created = time.Time{} // Handle parsing error (e.g., set to zero time)
		}

		updated, err := utils.ParseStringToTime(issue.IssueFields.Updated, format)
		if err != nil {
			updated = time.Time{} // Handle parsing error
		}

		dueDate, err := utils.ParseStringToTime(issue.IssueFields.DueDate, format)
		if err != nil {
			dueDate = time.Time{} // Handle parsing error
		}

		statusCategoryChanged, err := utils.ParseStringToTime(issue.IssueFields.StatusCategoryChangedDate, format)
		if err != nil {
			statusCategoryChanged = time.Time{} // Handle parsing error
		}

		if issue.Key == "BIT-21492" {
			log.Printf("%d - %d - %d - %d",
				issue.IssueFields.TimeOriginalEstimate,
				issue.IssueFields.AggregateTimeOriginalEstimate,
				issue.IssueFields.TimeEstimate,
				issue.IssueFields.AggregateTimeEstimate)
		}

		issues = append(issues, JiraIssue{
			Key:                           issue.Key,
			Self:                          issue.Self,
			AssigneeEmail:                 issue.IssueFields.Assignee.Email,
			AssigneeName:                  issue.IssueFields.Assignee.DisplayName,
			ReporterEmail:                 issue.IssueFields.Reporter.Email,
			ReporterName:                  issue.IssueFields.Reporter.DisplayName,
			CreatorEmail:                  issue.IssueFields.Reporter.Email,
			CreatorName:                   issue.IssueFields.Reporter.DisplayName,
			Description:                   issue.IssueFields.Description,
			Created:                       created,
			Updated:                       updated,
			DueDate:                       &dueDate,
			StatusCategoryChange:          &statusCategoryChanged,
			TimeOriginalEstimate:          utils.SafeFloat64(issue.IssueFields.TimeOriginalEstimate, 0),
			AggregateTimeOriginalEstimate: utils.SafeFloat64(issue.IssueFields.AggregateTimeOriginalEstimate, 0),
			TimeEstimate:                  utils.SafeFloat64(issue.IssueFields.TimeEstimate, 0),
			AggregateTimeEstimate:         utils.SafeFloat64(issue.IssueFields.AggregateTimeEstimate, 0),
			IssueTypeName:                 issue.IssueFields.IssueType.Name,
			IssueTypeDescription:          issue.IssueFields.IssueType.Description,
			ProjectKey:                    issue.IssueFields.Project.Key,
			ProjectName:                   issue.IssueFields.Project.Name,
		})
	}

	return issues
}
