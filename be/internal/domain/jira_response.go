package domain

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
	Resolution                    interface{}   `json:"resolution"`
	LastViewed                    interface{}   `json:"lastViewed"`
	Labels                        []string      `json:"labels"`
	AggregateTimeOriginalEstimate interface{}   `json:"aggregatetimeoriginalestimate"`
	IssueLinks                    []interface{} `json:"issuelinks"`
	Assignee                      JiraUser      `json:"assignee"`
	Components                    []interface{} `json:"components"`
	Subtasks                      []interface{} `json:"subtasks"`
	Customfield10164              CustomField   `json:"customfield_10164"`
	Reporter                      JiraUser      `json:"reporter"`
	Customfield10045              string        `json:"customfield_10045"`
	Customfield10046              float64       `json:"customfield_10046"`
	Progress                      Progress      `json:"progress"`
	IssueType                     IssueType     `json:"issuetype"`
	Project                       Project       `json:"project"`
	Customfield10032              []interface{} `json:"customfield_10032"`
	Customfield10157              []interface{} `json:"customfield_10157"`
	ResolutionDate                interface{}   `json:"resolutiondate"`
	Watches                       Watches       `json:"watches"`
	Customfield10019              string        `json:"customfield_10019"`
	Created                       string        `json:"created"`
	Updated                       string        `json:"updated"`
	TimeOriginalEstimate          interface{}   `json:"timeoriginalestimate"`
	Description                   string        `json:"description"`
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
