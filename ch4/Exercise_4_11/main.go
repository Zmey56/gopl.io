// Exercise 4.11:
// Build a tool that lets users create, read, update, and delete GitHub
// issues from the command line, invoking their preferred text editor
// when substantial text input is required.”

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	baseURL = "https://api.github.com/repos"
)

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	owner := "Zmey56"
	repo := "test"
	token := os.Getenv("YOUR_GITHUB_ACCESS_TOKEN")

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Create Issue")
		fmt.Println("2. Read Issue")
		fmt.Println("3. Update Issue")
		fmt.Println("4. Delete Issue")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createIssue(owner, repo, token)
		case 2:
			readIssue(owner, repo)
		case 3:
			updateIssue(owner, repo, token)
		case 4:
			deleteIssue(owner, repo, token)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func createIssue(owner, repo, token string) {
	title, body := getInputFromEditor()

	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, owner, repo)
	payload, err := json.Marshal(Issue{
		Title: title,
		Body:  body,
	})
	if err != nil {
		log.Fatalf("Failed to marshal JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Issue created successfully!")
	} else {
		fmt.Println("Failed to create issue. Please try again.")
	}
}

func readIssue(owner, repo string) {
	fmt.Print("Enter the issue number: ")
	var issueNumber int
	fmt.Scanln(&issueNumber)

	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, owner, repo, issueNumber)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}
		fmt.Println(string(body))
	} else {
		fmt.Println("Failed to get issue. Please try again.")
	}
}

func updateIssue(owner, repo, token string) {
	fmt.Print("Enter the issue number: ")
	var issueNumber int
	fmt.Scanln(&issueNumber)

	title, body := getInputFromEditor()

	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, owner, repo, issueNumber)
	payload, err := json.Marshal(Issue{
		Title: title,
		Body:  body,
	})
	if err != nil {
		log.Fatalf("Failed to marshal JSON payload: %v", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Issue updated successfully!")
	} else {
		fmt.Println("Failed to update issue. Please try again.")
	}
}

func deleteIssue(owner, repo, token string) {
	fmt.Print("Enter the issue number: ")
	var issueNumber int
	fmt.Scanln(&issueNumber)

	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, owner, repo, issueNumber)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Issue deleted successfully!")
	} else {
		fmt.Println("Failed to delete issue. Please try again.")
	}
}

func getInputFromEditor() (string, string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano" // Set a default editor (e.g., nano)
	}

	file, err := os.CreateTemp("", "issue_*.md")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	cmd := exec.Command(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to open editor: %v", err)
	}

	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.SplitN(string(content), "\n", 2)
	title := strings.TrimSpace(lines[0])
	body := ""
	if len(lines) > 1 {
		body = strings.TrimSpace(lines[1])
	}

	return title, body
}
