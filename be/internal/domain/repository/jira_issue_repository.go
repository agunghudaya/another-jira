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
func MapJiraResponseToJiraIssues(jiraResponse JiraResponse) []JiraIssue {
	var issues []JiraIssue

	format := "2006-01-02T15:04:05.000-0700"

	for _, issue := range jiraResponse.Issues {

		created, err := utils.ParseStringToTime(issue.Fields.Created, format)
		if err != nil {
			created = time.Time{} // Handle parsing error (e.g., set to zero time)
		}

		updated, err := utils.ParseStringToTime(issue.Fields.Updated, format)
		if err != nil {
			updated = time.Time{} // Handle parsing error
		}

		dueDate, err := utils.ParseStringToTime(issue.Fields.DueDate, format)
		if err != nil {
			dueDate = time.Time{} // Handle parsing error
		}

		statusCategoryChanged, err := utils.ParseStringToTime(issue.Fields.StatusCategoryChangedDate, format)
		if err != nil {
			statusCategoryChanged = time.Time{} // Handle parsing error
		}

		if issue.Key == "BIT-21492" {
			log.Printf("%d - %d - %d - %d",
				issue.Fields.TimeOriginalEstimate,
				issue.Fields.AggregateTimeOriginalEstimate,
				issue.Fields.TimeEstimate,
				issue.Fields.AggregateTimeEstimate)
		}

		issues = append(issues, JiraIssue{
			Key:                           issue.Key,
			Self:                          issue.Self,
			AssigneeEmail:                 issue.Fields.Assignee.Email,
			AssigneeName:                  issue.Fields.Assignee.DisplayName,
			ReporterEmail:                 issue.Fields.Reporter.Email,
			ReporterName:                  issue.Fields.Reporter.DisplayName,
			CreatorEmail:                  issue.Fields.Reporter.Email,
			CreatorName:                   issue.Fields.Reporter.DisplayName,
			Description:                   issue.Fields.Description,
			Created:                       created,
			Updated:                       updated,
			DueDate:                       &dueDate,
			StatusCategoryChange:          &statusCategoryChanged,
			TimeOriginalEstimate:          utils.SafeFloat64(issue.Fields.TimeOriginalEstimate, 0),
			AggregateTimeOriginalEstimate: utils.SafeFloat64(issue.Fields.AggregateTimeOriginalEstimate, 0),
			TimeEstimate:                  utils.SafeFloat64(issue.Fields.TimeEstimate, 0),
			AggregateTimeEstimate:         utils.SafeFloat64(issue.Fields.AggregateTimeEstimate, 0),
			IssueTypeName:                 issue.Fields.IssueType.Name,
			IssueTypeDescription:          issue.Fields.IssueType.Description,
			ProjectKey:                    issue.Fields.Project.Key,
			ProjectName:                   issue.Fields.Project.Name,
		})
	}

	return issues
}
