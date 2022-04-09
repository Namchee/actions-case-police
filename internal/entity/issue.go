package entity

import (
	"github.com/google/go-github/v43/github"
)

// IssueData represents fixed case GitHub issue
type IssueData struct {
	Title string
	Body  string
}

func GetIssueData(
	issue *github.PullRequest,
	dictionary map[string]string,
) *IssueData {
	return &IssueData{}
}
