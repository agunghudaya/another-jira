package jira

import (
	"be/internal/domain/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client represents a Jira API client
type Client struct {
	baseURL    string
	username   string
	apiToken   string
	httpClient *http.Client
}

// NewClient creates a new Jira client
func NewClient(cfg config.JiraConfig) (*Client, error) {
	client := &Client{
		baseURL:  cfg.BaseURL,
		username: cfg.Username,
		apiToken: cfg.APIToken,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}

	// Test the connection
	if err := client.testConnection(); err != nil {
		return nil, fmt.Errorf("failed to connect to Jira: %w", err)
	}

	return client, nil
}

// testConnection tests the connection to Jira
func (c *Client) testConnection() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/api/2/myself", c.baseURL), nil)
	if err != nil {
		return err
	}

	c.setAuth(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// setAuth sets the authentication headers
func (c *Client) setAuth(req *http.Request) {
	req.SetBasicAuth(c.username, c.apiToken)
}

// GetIssues retrieves issues from Jira
func (c *Client) GetIssues(jql string, startAt, maxResults int) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/api/2/search", c.baseURL)

	query := map[string]interface{}{
		"jql":        jql,
		"startAt":    startAt,
		"maxResults": maxResults,
		"fields": []string{
			"summary",
			"description",
			"status",
			"assignee",
			"reporter",
			"created",
			"updated",
			"duedate",
			"timeoriginalestimate",
			"timeestimate",
			"aggregatetimeoriginalestimate",
			"aggregatetimeestimate",
		},
	}

	body, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	c.setAuth(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result bytes.Buffer
	if _, err := result.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
