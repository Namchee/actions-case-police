package utils

import (
	"fmt"
	"regexp"

	"github.com/google/go-github/v43/github"
)

// IssueData represents fixed case GitHub issue
type IssueData struct {
	Title   string
	Body    string
	Changes map[string]string
}

// PolicizeIssue applies case police to issue title and body
func PolicizeIssue(
	issue *github.Issue,
	dictionary map[string]string,
) *IssueData {
	title := issue.GetTitle()
	body := issue.GetBody()

	changes := map[string]string{}

	for entry, replacer := range dictionary {
		regex := regexp.MustCompile(
			fmt.Sprintf(`(?i)\b%s\b`, entry),
		)

		replaced := false

		if regex.Match([]byte(title)) {
			title = regex.ReplaceAllString(title, replacer)
			replaced = true
		}

		if regex.Match([]byte(body)) {
			body = regex.ReplaceAllString(body, replacer)
			replaced = true
		}

		if replaced {
			changes[entry] = replacer
		}
	}

	return &IssueData{
		Title:   title,
		Body:    body,
		Changes: changes,
	}
}
