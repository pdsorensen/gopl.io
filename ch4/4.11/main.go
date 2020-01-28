package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// get access token and insert where relevant

func main() {
	command := os.Args[1]

	switch command {
	case "create":
		createIssue(os.Args[2], os.Args[3], os.Args[4])
	case "list":
		listIssues(os.Args[2])
	case "get":
		getIssue(os.Args[2], os.Args[3])
	case "close":
		editIssue(os.Args[2], os.Args[3], "close", nil)
	case "edit":
		fields := prepareEdit(os.Args[2], os.Args[3])
		editIssue(os.Args[2], os.Args[3], "open", fields)
	}
}

func createIssue(repo string, title string, body string) {
	values := map[string]string{"title": title, "body": body}
	jsonValue, _ := json.Marshal(values)
	url := fmt.Sprintf("%s/repos/%s/%s/issues", APIURL, "pdsorensen", repo)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "token TOKEN")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
}

func getIssue(repo string, issueID string) (*Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", APIURL, "pdsorensen", repo, issueID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token TOKEN")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func listIssues(repo string) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues", APIURL, "pdsorensen", repo)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token TOKEN")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("Issue was successfully created")
}

func prepareEdit(repo string, issueID string) map[string]string {
	editor := "vim"
	editorPath, err := exec.LookPath(editor)

	tempfile, _ := ioutil.TempFile("", "issue_crud")
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, _ := getIssue(repo, issueID)
	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})

	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	return fields
}

func editIssue(repo string, issueID string, state string, body map[string]string) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", APIURL, "pdsorensen", repo, issueID)

	values := map[string]string{"state": state}

	if body != nil {
		values = body
	}

	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "token 7ad9903054b843519944a3a6b1ecae76afdae195")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	htmlData, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(htmlData))
}
