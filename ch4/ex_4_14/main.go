//“Exercise 4.14:
//Create a web server that queries GitHub once and then allows
//navigation of the list of bug reports, milestones, and users.”

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var GITHUB_URL string = "https://api.github.com"

type UsersResult struct {
	Users []*User
}

type User struct {
	Id    int
	Login string
}

type RepositoriesResult struct {
	Collection []Repositories
}
type Repositories struct {
	Id   int
	Name string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/repositories", repositories)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := getUsers()
	if err != nil {
		log.Printf("Error: %v", err.Error())
		html := "<html>"
		html += "<head>"
		html += "</head>"
		html += "<body>"
		html += "Error!"
		html += "</body>"
		html += "</html>"

		fmt.Fprintf(w, "%s\n", html)
	} else {
		html := "<html>"
		html += "<head>"
		html += "</head>"
		html += "<body>"
		html += "<h1>Users</h1>"
		for i := range *res {
			html += "<p>"
			html += "Repositories for: "
			html += "<a href=\"/repositories?u=" + (*res)[i].Login + "\">"
			html += (*res)[i].Login + "</a>"
			html += "</p>"
		}
		html += "</body>"
		html += "</html>"
		fmt.Fprintf(w, "%s\n", html)
	}
}

func repositories(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query()["u"]
	if user != nil {
		res, err := getRepositories(user[0])
		if err != nil {
			html := "<html>"
			html += "<head>"
			html += "</head>"
			html += "<body>"
			html += "No such user."
			html += "</body>"
			html += "</html>"
			fmt.Fprintf(w, "%s\n", html)
			log.Printf("%v", err)
		} else {
			html := "<html>"
			html += "<head>"
			html += "</head>"
			html += "<body>"
			html += "<h1>Users</h1>"
			for i := range *res {
				html += "<p>"
				html += "User: " + user[0] + " Repository: " + (*res)[i].Name
				html += "</p>"
			}
			html += "</body>"
			html += "</html>"
			fmt.Fprintf(w, "%s\n", html)
		}

	} else {
		html := "<html>"
		html += "<head>"
		html += "</head>"
		html += "<body>"
		html += "No such user."
		html += "</body>"
		html += "</html>"
		fmt.Fprintf(w, "%s\n", html)
	}
}

func getRepositories(user string) (*[]Repositories, error) {
	q := url.QueryEscape(user)
	url := GITHUB_URL + "/users/" + q + "/repos"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	repos := make([]Repositories, 0)

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return &repos, nil
}

func getUsers() (*[]User, error) {
	resp, err := http.Get(GITHUB_URL + "/users")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var users []User

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return &users, nil
}
