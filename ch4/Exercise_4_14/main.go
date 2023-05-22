//Exercise 4.14
//Create a web server that queries GitHub once and then allows
//navigation of the list of bug reports, milestones, and users.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Replace the placeholders below with your own GitHub credentials.
const (
	GitHubOwner = "golang"
	GitHubRepo  = "go"
)

type Issue struct {
	Title string `json:"title"`
	URL   string `json:"html_url"`
}

type Milestone struct {
	Title string `json:"title"`
	URL   string `json:"html_url"`
}

type User struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Contributions     int    `json:"contributions"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		issues, err := getIssues()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		milestones, err := getMilestones()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users, err := getUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Bug Reports:\n")
		for _, issue := range issues {
			fmt.Fprintf(w, "- %s (%s)\n", issue.Title, issue.URL)
		}

		fmt.Fprintf(w, "\nMilestones:\n")
		for _, milestone := range milestones {
			fmt.Fprintf(w, "- %s (%s)\n", milestone.Title, milestone.URL)
		}

		fmt.Fprintf(w, "\nUsers:\n")
		for _, user := range users {
			fmt.Fprintf(w, "- %s (%s)\n", user.Login, user.Url)
		}
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getIssues() ([]Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", GitHubOwner, GitHubRepo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func getMilestones() ([]Milestone, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/milestones", GitHubOwner, GitHubRepo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var milestones []Milestone
	err = json.NewDecoder(resp.Body).Decode(&milestones)
	if err != nil {
		return nil, err
	}

	return milestones, nil
}

func getUsers() ([]User, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors", GitHubOwner, GitHubRepo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
