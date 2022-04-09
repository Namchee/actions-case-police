package service

import (
	"testing"

	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/google/go-github/v43/github"
	"github.com/stretchr/testify/assert"
)

func TestPolicizeIssue(t *testing.T) {
	sampleTitle := "github gitlab vscode"
	sampleBodyReplace := `## Overview
	
	Lorem ipsum dolor sil GITHUB Gitlab vScOdE.`
	sampleBodyNoReplace := `## Overview
	
	Lorem ipsum dolor sil githubgitlabvscode.`

	tests := []struct {
		name string
		args *github.Issue
		cfg  *entity.Configuration
		want *entity.IssueData
	}{
		{
			name: "should replace all entries",
			args: &github.Issue{
				Title: &sampleTitle,
				Body:  &sampleBodyReplace,
			},
			cfg: &entity.Configuration{
				Dictionary: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
				Fix: true,
			},
			want: &entity.IssueData{
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
			cfg: &entity.Configuration{
				Dictionary: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
				Fix: true,
			},
			want: &entity.IssueData{
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
			cfg: &entity.Configuration{
				Dictionary: map[string]string{},
				Fix:        true,
			},
			want: &entity.IssueData{
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
			cfg: &entity.Configuration{
				Dictionary: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
				Fix: true,
			},
			want: &entity.IssueData{
				Title:   "",
				Body:    "",
				Changes: map[string]string{},
			},
		},
		{
			name: "should not anything as the contents are empty",
			args: &github.Issue{
				Title: nil,
				Body:  nil,
			},
			cfg: &entity.Configuration{
				Dictionary: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
				Fix: true,
			},
			want: &entity.IssueData{
				Title:   "",
				Body:    "",
				Changes: map[string]string{},
			},
		},
		{
			name: "should not do anything as the contents are empty",
			args: &github.Issue{
				Title: nil,
				Body:  nil,
			},
			cfg: &entity.Configuration{
				Dictionary: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
				Fix: true,
			},
			want: &entity.IssueData{
				Title:   "",
				Body:    "",
				Changes: map[string]string{},
			},
		},
		{
			name: "should not do anything as fix is false",
			args: &github.Issue{
				Title: &sampleTitle,
				Body:  &sampleBodyReplace,
			},
			cfg: &entity.Configuration{
				Dictionary: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
				Fix: false,
			},
			want: &entity.IssueData{
				Title: sampleTitle,
				Body:  sampleBodyReplace,
				Changes: map[string]string{
					"github": "GitHub",
					"gitlab": "GitLab",
					"vscode": "VS Code",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PolicizeIssue(tc.args, tc.cfg)

			assert.Equal(t, tc.want, got)
		})
	}
}
