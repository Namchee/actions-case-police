package service

import (
	"fmt"
	"regexp"

	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/google/go-github/v43/github"
)

// PolicizeIssue applies case police to issue title and body
func PolicizeIssue(
	issue *github.Issue,
	cfg *entity.Configuration,
) *entity.IssueData {
	title := issue.GetTitle()
	body := issue.GetBody()

	changes := map[string]string{}

	for entry, replacer := range cfg.Dictionary {
		regex := regexp.MustCompile(
			fmt.Sprintf(`(?i)\b%s\b`, entry),
		)

		found := false

		if regex.Match([]byte(title)) || regex.Match([]byte(body)) {
			found = true

			if cfg.Fix {
				body = regex.ReplaceAllString(body, replacer)
				title = regex.ReplaceAllString(title, replacer)
			}
		}

		if found {
			changes[entry] = replacer
		}
	}

	return &entity.IssueData{
		Title:   title,
		Body:    body,
		Changes: changes,
	}
}
