package jira_atlassian

import (
	domain "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (r *jiraAtlassianRepository) FetchJiraTasksWithFilter(ctx context.Context, jiraUserID string, cfg *config.Config) (jiraResp domain.JiraIssueResponse, err error) {

	jiraData := domain.JiraIssueResponse{}
	startAt := 0

	jql := fmt.Sprintf(`
	assignee = %s 
	and status not in (CANCELLED, OPEN) 
	and issuetype != Epic 
	and (created >= "2025-01-01" or resolved >= "2025-01-01") 
	order by status DESC, created DESC`, jiraUserID)

	client := &http.Client{}
	reqURL := fmt.Sprintf("%s%s?jql=%s&&maxResults=50&startAt=%d", r.cfg.GetString("jira.baseurl"), r.cfg.GetString("jira.searchurl"), url.QueryEscape(jql), startAt)

	r.log.Printf("url: %s", reqURL)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return jiraData, err
	}

	req.SetBasicAuth(cfg.GetString("jira_username"), cfg.GetString("jira_token"))
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return jiraData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return jiraData, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&jiraResp)
	if err != nil {
		return jiraData, err
	}

	return jiraResp, nil
}
