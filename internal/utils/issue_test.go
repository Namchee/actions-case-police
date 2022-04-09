package utils

import (
	"testing"

	"github.com/google/go-github/v43/github"
	"github.com/stretchr/testify/assert"
)

func TestPolicizeIssue(t *testing.T) {
	sampleTitle := "github gitlab vscode"
	sampleBodyReplace := `## Overview
	
	Lorem ipsum dolor sil github gitlab vscode.`
	sampleBodyNoReplace := `## Overview
	
	Lorem ipsum dolor sil githubgitlabvscode.`

	tests := []struct {
		name string
		args *github.Issue
		dict map[string]string
		want *IssueData
	}{
		{
			name: "should replace all entries",
			args: &github.Issue{
				Title: &sampleTitle,
				Body:  &sampleBodyReplace,
			},
			dict: map[string]string{
				"github": "GitHub",
				"gitlab": "GitLab",
				"vscode": "VS Code",
			},
			want: &IssueData{
				Title: "GitHub GitLab VS Code",
				Body: `## Overview
	
	Lorem ipsum dolor sil GitHub GitLab VS Code.`,
				Changes: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
			},
		},
		{
			name: "should replace nothing in the body as there is no word boundary",
			args: &github.Issue{
				Title: &sampleTitle,
				Body:  &sampleBodyNoReplace,
			},
			dict: map[string]string{
				"github": "GitHub",
				"gitlab": "GitLab",
				"vscode": "VS Code",
			},
			want: &IssueData{
				Title: "GitHub GitLab VS Code",
				Body:  sampleBodyNoReplace,
				Changes: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
			},
		},
		{
			name: "should replace nothing as dictionary is empty",
			args: &github.Issue{
				Title: &sampleTitle,
				Body:  &sampleBodyReplace,
			},
			dict: map[string]string{},
			want: &IssueData{
				Title:   sampleTitle,
				Body:    sampleBodyReplace,
				Changes: map[string]string{},
			},
		},
		{
			name: "should not anything as the contents are empty",
			args: &github.Issue{
				Title: nil,
				Body:  nil,
			},
			dict: map[string]string{
				"github": "GitHub",
				"gitlab": "GitLab",
				"vscode": "VS Code",
			},
			want: &IssueData{
				Title:   "",
				Body:    "",
				Changes: map[string]string{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PolicizeIssue(tc.args, tc.dict)

			assert.Equal(t, tc.want, got)
		})
	}
}
