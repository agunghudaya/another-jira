// internal/usecase/sync_service.go
package usecase

import (
	"be/internal/domain"
	"be/internal/infrastructure/config"
	"be/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type JiraSync interface {
	ProcessSync() error
	JiraUserSync(user *domain.User) error
}

type jiraSync struct {
	cfg      *config.Config
	syncRepo repository.SyncRepository
}

func NewJiraSyncUsecase(cfg *config.Config, syncRepo repository.SyncRepository) JiraSync {
	return &jiraSync{cfg: cfg, syncRepo: syncRepo}
}

func (s *jiraSync) ProcessSync() error {
	// Fetch pending sync
	users, err := s.syncRepo.FetchUserList()
	if err != nil {
		return err
	}

	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return err
	}

	if err != nil || len(users) == 0 {
		log.Println("No user found. Err:", err)
		return nil
	}

	for _, user := range users {
		s.JiraUserSync(&user)
	}

	return nil
}

func (s *jiraSync) JiraUserSync(user *domain.User) error {

	log.Println(user.JiraUserID)
	syncHistory, err := s.syncRepo.FetchPendingSync(user.JiraUserID)

	if err != nil {
		log.Println("FetchUserList fail with err:", err)
		return err
	}

	log.Println("syncHistory", syncHistory)

	return nil
}

func (s *jiraSync) FetchJiraTasksWithFilter() (domain.JiraResponse, error) {

	jiraData := domain.JiraResponse{}
	startAt := 0

	jql := fmt.Sprintf(`
	assignee = %s 
	and status not in (CANCELLED, OPEN) 
	and issuetype != Epic 
	and (created >= "2024-07-01" or resolved >= "2024-07-01") 
	order by status DESC, created DESC`, "6277ef5a0e2c49006901f266")

	client := &http.Client{}
	reqURL := fmt.Sprintf("%s%s?jql=%s&&maxResults=50&startAt=%d", s.cfg.GetString("jira.baseurl"), s.cfg.GetString("jira.searchurl"), url.QueryEscape(jql), startAt)

	fmt.Println(reqURL)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return jiraData, err
	}

	req.SetBasicAuth(s.cfg.GetString("JIRA_USERNAME"), s.cfg.GetString("JIRA_TOKEN"))
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return jiraData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return jiraData, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	var tmpJiraData *domain.JiraResponse

	err = json.NewDecoder(resp.Body).Decode(&tmpJiraData)
	if err != nil {
		return jiraData, err
	}

	err = json.NewDecoder(resp.Body).Decode(&tmpJiraData)
	if err != nil {
		return jiraData, err
	}

	jiraData.Issues = append(jiraData.Issues, tmpJiraData.Issues...)

	return jiraData, nil
}
