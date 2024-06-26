package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func FetchProjects(apiKey string, skills []string) ([]Project, error) {
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
	// q.Add("min_avg_price", "30")
	q.Add("sort_field", "time_updated")
	q.Add("limit", "100")
	q.Add("query", strings.Join(skills, " "))

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ResponseProjects
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.Result.TotalCount)

	if response.Status != "success" {
		return nil, fmt.Errorf("Error fetching projects!")
	}

	return response.Result.Projects, nil
}

func FetchOwners(apiKey string, ownerIDs []int) ([]Owner, error) {
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

	for _, ID := range ownerIDs {
		q.Add("users[]", strconv.Itoa(ID))
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println((string(body))[:5000])

	var response ResponseUsers
	err = json.Unmarshal(body, &response)
	fmt.Println(response.Status)
	if err != nil {
		return nil, err
	}
	var owners []Owner
	for _, owner := range response.Result.Users {
		owners = append(owners, owner)
	}

	return owners, nil
}

func CountCommonIDs(arr1, arr2 []Project) int {
	idMap := make(map[int]bool)
	for _, proj := range arr1 {
		idMap[proj.ID] = true
	}

	commonCount := 0
	for _, proj := range arr2 {
		if idMap[proj.ID] {
			commonCount++
		}
	}
	return commonCount
}
