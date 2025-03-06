package domain

type JiraResponse struct {
	Issues []struct {
		Key    string `json:"key"`
		Fields Fields `json:"fields"`
	} `json:"issues"`
}

type Fields struct {
	Summary                   string    `json:"summary"`
	Parent                    *Parent   `json:"parent"` // Handle parent field
	Status                    Status    `json:"status"`
	Assignee                  Assignee  `json:"assignee"`
	Sprint                    *[]Sprint `json:"customfield_10020"` // Handle array of Sprint or null
	StartDate                 string    `json:"customfield_10015"`
	EndDate                   string    `json:"customfield_10051"`
	UpdateTime                string    `json:"updated"`
	CreateTime                string    `json:"created"`
	ResolveTime               string    `json:"resolutiondate"`
	Creator                   *Creator  `json:"creator"`
	DueDate                   string    `json:"duedate"`
	StatusCategoryChangedDate string    `json:"statuscategorychangedate"`
	IssueType                 IssueType `json:"issuetype"`
}

type Parent struct {
	ID     string       `json:"id"`
	Key    string       `json:"key"`
	Fields ParentFields `json:"fields"`
}

type ParentFields struct {
	Summary string       `json:"summary"`
	Status  ParentStatus `json:"status"`
}

type ParentStatus struct {
	Name string `json:"name"`
}

type Status struct {
	Name string `json:"name"`
}

type Assignee struct {
	DisplayName string `json:"displayName"`
}

type Sprint struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	State        string `json:"state"`
	BoardID      int    `json:"boardId"`
	Goal         string `json:"goal"`
	StartDate    string `json:"startDate,omitempty"`
	EndDate      string `json:"endDate,omitempty"`
	CompleteDate string `json:"completeDate,omitempty"`
}

type IssueType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Creator struct {
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
}
