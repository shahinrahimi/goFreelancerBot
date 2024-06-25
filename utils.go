package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchProjects(apiKey string) ([]Project, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.freelancer.com/api/projects/0.1/projects/active/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("freelancer-oauth-v1", apiKey)
	q := req.URL.Query()
	q.Add("project_types", "fixed")
	q.Add("from_time", "1")
	q.Add("offset", "0")
	q.Add("frontend_project_statuses", "open")
	q.Add("full_description", "true")
	q.Add("attachment_details", "true")
	q.Add("or_search_query", "true")
	q.Add("min_avg_price", "30")
	q.Add("sort_field", "time_updated")
	q.Add("limit", "100")
	q.Add("query", strings.Join([]string{
		"python",
		"javascript",
		"excel",
		"vba",
		"libreoffice",
		"shiny",
		"scraping",
		"react",
		"go",
	}, " "))

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ResponseProjects
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("Error fetching projects!")
	}

	return response.Result.Projects, nil
}

func FetchOwners(apiKey string, ownerIDs []string) ([]Owner, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.freelancer.com/api/users/0.1/users/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("freelancer-oauth-v1", apiKey)

	q := req.URL.Query()
	q.Add("display_info", "true")
	q.Add("status", "true")
	q.Add("reputation", "true")
	q.Add("employer_reputation", "true")
	q.Add("country_details", "true")
	q.Add("avatar", "true")

	ownerIDsStr := make([]string, len(ownerIDs))
	for i, id := range ownerIDs {
		ownerIDsStr[i] = fmt.Sprint(id)
	}
	q.Add("users", strings.Join(ownerIDsStr, ","))

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Owners []Owner `json:"owners"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Owners, nil
}
