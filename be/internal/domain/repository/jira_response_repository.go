package repository

type JiraResponse struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Assignee                  JiraUser      `json:"assignee"`
	Components                []interface{} `json:"components"`
	Created                   string        `json:"created"`
	Customfield10019          string        `json:"customfield_10019"`
	Customfield10032          []interface{} `json:"customfield_10032"`
	Customfield10045          string        `json:"customfield_10045"`
	Customfield10046          float64       `json:"customfield_10046"`
	Customfield10157          []interface{} `json:"customfield_10157"`
	Customfield10164          CustomField   `json:"customfield_10164"`
	Description               string        `json:"description"`
	DueDate                   string        `json:"duedate"`
	IssueLinks                []interface{} `json:"issuelinks"`
	IssueType                 IssueType     `json:"issuetype"`
	Labels                    []string      `json:"labels"`
	LastViewed                interface{}   `json:"lastViewed"`
	Progress                  Progress      `json:"progress"`
	Project                   Project       `json:"project"`
	Reporter                  JiraUser      `json:"reporter"`
	Resolution                interface{}   `json:"resolution"`
	ResolutionDate            interface{}   `json:"resolutiondate"`
	Subtasks                  []interface{} `json:"subtasks"`
	Updated                   string        `json:"updated"`
	Watches                   Watches       `json:"watches"`
	StatusCategoryChangedDate string        `json:"statuscategorychangedate"`

	//time estimation
	AggregateTimeEstimate         interface{} `json:"aggregatetimeestimate"`
	AggregateTimeOriginalEstimate interface{} `json:"aggregatetimeoriginalestimate"`
	TimeEstimate                  interface{} `json:"timeestimate"`
	TimeOriginalEstimate          interface{} `json:"timeoriginalestimate"`
}

type JiraUser struct {
	Self        string            `json:"self"`
	AccountID   string            `json:"accountId"`
	Email       string            `json:"emailAddress"`
	AvatarUrls  map[string]string `json:"avatarUrls"`
	DisplayName string            `json:"displayName"`
	Active      bool              `json:"active"`
	TimeZone    string            `json:"timeZone"`
	AccountType string            `json:"accountType"`
}

type CustomField struct {
	Self  string `json:"self"`
	Value string `json:"value"`
	ID    string `json:"id"`
}

type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
	AvatarID    int    `json:"avatarId"`
	Hierarchy   int    `json:"hierarchyLevel"`
}

type Project struct {
	Self            string            `json:"self"`
	ID              string            `json:"id"`
	Key             string            `json:"key"`
	Name            string            `json:"name"`
	ProjectTypeKey  string            `json:"projectTypeKey"`
	Simplified      bool              `json:"simplified"`
	AvatarUrls      map[string]string `json:"avatarUrls"`
	ProjectCategory ProjectCategory   `json:"projectCategory"`
}

type ProjectCategory struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Watches struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

type CustomField18 struct {
	HasEpicLinkFieldDependency bool              `json:"hasEpicLinkFieldDependency"`
	ShowField                  bool              `json:"showField"`
	NonEditableReason          NonEditableReason `json:"nonEditableReason"`
}

type NonEditableReason struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}
