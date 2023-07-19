package policy

import (
	"context"
	"digger/pkg/ci"
	"digger/pkg/core/policy"
	"errors"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"io"
	"log"
	"net/http"
	"net/url"
)

type DiggerHttpPolicyProvider struct {
	DiggerHost         string
	DiggerOrganisation string
	AuthToken          string
	HttpClient         *http.Client
}

type NoOpPolicyChecker struct {
}

func (p NoOpPolicyChecker) CheckAccessPolicy(_ ci.OrgService, _ string, _ string, _ string, _ string, _ string) (bool, error) {
	return true, nil
}

func (p NoOpPolicyChecker) CheckPlanPolicy(_ string, _ string, _ string) (bool, error) {
	return true, nil
}

func getPolicyForOrganisation(p *DiggerHttpPolicyProvider) (string, *http.Response, error) {
	organisation := p.DiggerOrganisation
	u, err := url.Parse(p.DiggerHost)
	if err != nil {
		log.Fatalf("Not able to parse digger cloud url: %v", err)
	}
	u.Path = "/orgs/" + organisation + "/access-policy"
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Authorization", "Bearer "+p.AuthToken)

	resp, err := p.HttpClient.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp, nil
	}
	return string(body), resp, nil
}

func getPolicyForNamespace(p *DiggerHttpPolicyProvider, namespace string, projectName string) (string, *http.Response, error) {
	// fetch RBAC policies for project from Digger API
	u, err := url.Parse(p.DiggerHost)
	if err != nil {
		log.Fatalf("Not able to parse digger cloud url: %v", err)
	}
	u.Path = "/repos/" + namespace + "/projects/" + projectName + "/access-policy"
	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Authorization", "Bearer "+p.AuthToken)

	resp, err := p.HttpClient.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp, nil
	}
	return string(body), resp, nil

}

// GetPolicy fetches policy for particular project,  if not found then it will fallback to org level policy
func (p *DiggerHttpPolicyProvider) GetAccessPolicy(organisation string, repo string, projectName string) (string, error) {
	namespace := fmt.Sprintf("%v-%v", organisation, repo)
	content, resp, err := getPolicyForNamespace(p, namespace, projectName)
	if err != nil {
		return "", err
	}

	// project policy found
	if resp.StatusCode == 200 && content != "" {
		return content, nil
	}

	// check if project policy was empty or not found (retrieve org policy if so)
	if (resp.StatusCode == 200 && content == "") || resp.StatusCode == 404 {
		content, resp, err := getPolicyForOrganisation(p)
		if err != nil {
			return "", err
		}
		if resp.StatusCode == 200 {
			return content, nil
		} else if resp.StatusCode == 404 {
			return "", nil
		} else {
			return "", errors.New(fmt.Sprintf("unexpected response while fetching organisation policy: %v, code %v", content, resp.StatusCode))
		}
	} else {
		return "", errors.New(fmt.Sprintf("unexpected response while fetching project policy: %v code %v", content, resp.StatusCode))
	}
}

func (p *DiggerHttpPolicyProvider) GetPlanPolicy(organisation string, repo string, projectName string) (string, error) {
	planPolicy := "package digger\n\ndeny[sprintf(message, [resource.address])] {\n  message := \"Cannot create EC2 instances!\"\n  resource := input.terraform.resource_changes[_]\n  resource.change.actions[_] == \"create\"\n  resource[type] == \"aws_instance\"\n}\n"
	return planPolicy, nil
}

func (p *DiggerHttpPolicyProvider) GetOrganisation() string {
	return p.DiggerOrganisation
}

type DiggerPolicyChecker struct {
	PolicyProvider policy.Provider
}

func (p DiggerPolicyChecker) CheckAccessPolicy(ciService ci.OrgService, SCMOrganisation string, SCMrepository string, projectName string, command string, requestedBy string) (bool, error) {
	// TODO: Get rid of organisation if its not needed
	organisation := p.PolicyProvider.GetOrganisation()
	policy, err := p.PolicyProvider.GetAccessPolicy(organisation, SCMrepository, projectName)

	if err != nil {
		fmt.Printf("Error while fetching policy: %v", err)
		return false, err
	}

	teams, err := ciService.GetUserTeams(SCMOrganisation, requestedBy)
	if err != nil {
		fmt.Printf("Error while fetching user teams for CI service: %v", err)
		return false, err
	}

	input := map[string]interface{}{
		"user":         requestedBy,
		"organisation": SCMOrganisation,
		"teams":        teams,
		"action":       command,
		"project":      projectName,
	}

	if policy == "" {
		return true, nil
	}

	ctx := context.Background()
	fmt.Printf("DEBUG: passing the following input policy: %v ||| text: %v", input, policy)
	query, err := rego.New(
		rego.Query("data.digger.allow"),
		rego.Module("digger", policy),
	).PrepareForEval(ctx)

	if err != nil {
		return false, err
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if len(results) == 0 || len(results[0].Expressions) == 0 {
		return false, fmt.Errorf("no result found")
	}

	expressions := results[0].Expressions

	for _, expression := range expressions {
		decision, ok := expression.Value.(bool)
		if !ok {
			return false, fmt.Errorf("decision is not a boolean")
		}
		if !decision {
			return false, nil
		}
	}

	return true, nil
}

func (p DiggerPolicyChecker) CheckPlanPolicy(SCMrepository string, projectName string, planOutput string) (bool, error) {
	// TODO: Get rid of organisation if its not needed
	organisation := p.PolicyProvider.GetOrganisation()
	policy, err := p.PolicyProvider.GetPlanPolicy(organisation, SCMrepository, projectName)

	input := map[string]interface{}{
		"terraform": planOutput,
	}

	if policy == "" {
		return true, nil
	}

	ctx := context.Background()
	fmt.Printf("DEBUG: passing the following input policy: %v ||| text: %v", input, policy)
	query, err := rego.New(
		rego.Query("data.digger.deny"),
		rego.Module("digger", policy),
	).PrepareForEval(ctx)

	if err != nil {
		return false, err
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if len(results) == 0 || len(results[0].Expressions) == 0 {
		return false, fmt.Errorf("no result found")
	}

	expressions := results[0].Expressions

	for _, expression := range expressions {
		decision, ok := expression.Value.([]string)
		if !ok {
			return false, fmt.Errorf("decision is not a boolean")
		}
		if len(decision) > 0 {
			for _, d := range decision {
				fmt.Printf("denied: %v\n", d)
			}
			return false, nil
		}
	}

	return true, nil
}
