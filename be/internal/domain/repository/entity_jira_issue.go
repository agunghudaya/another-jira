package repository

import (
	"be/internal/utils"
	"time"
)

type JiraIssueEntity struct {
	Key                  string
	AssigneeEmail        string
	AssigneeName         string
	CreatorEmail         string
	CreatorName          string
	Description          string
	IssueTypeDescription string
	IssueTypeName        string
	PriorityName         string
	ProjectID            int
	ProjectKey           string
	ProjectName          string
	ReporterEmail        string
	ReporterName         string
	Self                 string
	StatusCategoryKey    string
	StatusCategoryName   string
	StatusDescription    string
	StatusName           string
	Summary              string
	URL                  string

	// time tracking
	Created              time.Time
	Updated              time.Time
	DueDate              *time.Time
	StatusCategoryChange *time.Time

	//time estimation
	AggregateTimeEstimate         float64
	AggregateTimeOriginalEstimate float64
	TimeEstimate                  float64
	TimeOriginalEstimate          float64
}

func MapJiraResponseToJiraIssues(jiraResponse JiraIssueResponse) []JiraIssueEntity {
	var issues []JiraIssueEntity

	format := "2006-01-02T15:04:05.000-0700"
	formatShort := "2006-01-02"

	for _, issue := range jiraResponse.Issues {

		created, err := utils.ParseStringToTime(issue.IssueFields.Created, format)
		if err != nil {
			created = time.Time{} // Handle parsing error (e.g., set to zero time)
		}

		updated, err := utils.ParseStringToTime(issue.IssueFields.Updated, format)
		if err != nil {
			updated = time.Time{} // Handle parsing error
		}

		dueDate, err := utils.ParseStringToTime(issue.IssueFields.DueDate, formatShort)
		if err != nil {
			dueDate = time.Time{} // Handle parsing error
		}

		statusCategoryChanged, err := utils.ParseStringToTime(issue.IssueFields.StatusCategoryChangedDate, format)
		if err != nil {
			statusCategoryChanged = time.Time{} // Handle parsing error
		}

		issues = append(issues, JiraIssueEntity{
			AggregateTimeEstimate:         utils.SafeFloat64(issue.IssueFields.AggregateTimeEstimate, 0),
			AggregateTimeOriginalEstimate: utils.SafeFloat64(issue.IssueFields.AggregateTimeOriginalEstimate, 0),
			AssigneeEmail:                 issue.IssueFields.Assignee.Email,
			AssigneeName:                  issue.IssueFields.Assignee.DisplayName,
			Created:                       created,
			CreatorEmail:                  issue.IssueFields.Reporter.Email,
			CreatorName:                   issue.IssueFields.Reporter.DisplayName,
			Description:                   issue.IssueFields.Description,
			DueDate:                       &dueDate,
			IssueTypeDescription:          issue.IssueFields.IssueType.Description,
			IssueTypeName:                 issue.IssueFields.IssueType.Name,
			Key:                           issue.Key,
			PriorityName:                  issue.IssueFields.Priority.Name,
			ProjectKey:                    issue.IssueFields.Project.Key,
			ProjectName:                   issue.IssueFields.Project.Name,
			ReporterEmail:                 issue.IssueFields.Reporter.Email,
			ReporterName:                  issue.IssueFields.Reporter.DisplayName,
			Self:                          issue.Self,
			StatusCategoryChange:          &statusCategoryChanged,
			StatusCategoryKey:             issue.IssueFields.Status.StatusCategory.Key,
			StatusCategoryName:            issue.IssueFields.Status.StatusCategory.Name,
			StatusDescription:             issue.IssueFields.Status.Description,
			StatusName:                    issue.IssueFields.Status.Name,
			Summary:                       issue.IssueFields.Summary,
			TimeEstimate:                  utils.SafeFloat64(issue.IssueFields.TimeEstimate, 0),
			TimeOriginalEstimate:          utils.SafeFloat64(issue.IssueFields.TimeOriginalEstimate, 0),
			Updated:                       updated,
		})
	}

	return issues
}
