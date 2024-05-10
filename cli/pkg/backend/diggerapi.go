package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/diggerhq/digger/libs/orchestrator/scheduler"
	"github.com/diggerhq/digger/libs/terraform_utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

type NoopApi struct {
}

func (n NoopApi) ReportProject(namespace string, projectName string, configurationYaml string) error {
	return nil
}

func (n NoopApi) ReportProjectRun(namespace string, projectName string, startedAt time.Time, endedAt time.Time, status string, command string, output string) error {
	return nil
}

func (n NoopApi) ReportProjectJobStatus(repo string, projectName string, jobId string, status string, timestamp time.Time, summary *terraform_utils.PlanSummary) (*scheduler.SerializedBatch, error) {
	return nil, nil
}

type DiggerApi struct {
	DiggerHost string
	AuthToken  string
	HttpClient *http.Client
}

func (d DiggerApi) ReportProject(namespace string, projectName string, configurationYaml string) error {
	u, err := url.Parse(d.DiggerHost)
	if err != nil {
		log.Fatalf("Not able to parse digger cloud url: %v", err)
	}
	u.Path = filepath.Join(u.Path, "repos", namespace, "report-projects")

	request := map[string]interface{}{
		"name":              projectName,
		"configurationYaml": configurationYaml,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Not able to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("error while creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.AuthToken))

	resp, err := d.HttpClient.Do(req)

	if err != nil {
		return fmt.Errorf("error while sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status when reporting a project: %v", resp.StatusCode)
	}

	return nil
}

func (d DiggerApi) ReportProjectRun(namespace string, projectName string, startedAt time.Time, endedAt time.Time, status string, command string, output string) error {
	u, err := url.Parse(d.DiggerHost)
	if err != nil {
		log.Fatalf("Not able to parse digger cloud url: %v", err)
	}

	u.Path = filepath.Join(u.Path, "repos", namespace, "projects", projectName, "runs")

	request := map[string]interface{}{
		"startedAt": startedAt,
		"endedAt":   endedAt,
		"status":    status,
		"command":   command,
		"output":    output,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Not able to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("error while creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.AuthToken))

	resp, err := d.HttpClient.Do(req)

	if err != nil {
		return fmt.Errorf("error while sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status when reporting a project run: %v", resp.StatusCode)
	}

	return nil
}

func (d DiggerApi) ReportProjectJobStatus(repo string, projectName string, jobId string, status string, timestamp time.Time, summary *terraform_utils.PlanSummary) (*scheduler.SerializedBatch, error) {
	u, err := url.Parse(d.DiggerHost)
	if err != nil {
		log.Fatalf("Not able to parse digger cloud url: %v", err)
	}

	u.Path = filepath.Join(u.Path, "repos", repo, "projects", projectName, "jobs", jobId, "set-status")
	request := map[string]interface{}{
		"status":      status,
		"timestamp":   timestamp,
		"job_summary": summary.ToJson(),
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Not able to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, fmt.Errorf("error while creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.AuthToken))

	resp, err := d.HttpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error while sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status when reporting a project job status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response body: %v", err)
	}

	var response scheduler.SerializedBatch
	json.Unmarshal(body, &response)

	return &response, nil
}
