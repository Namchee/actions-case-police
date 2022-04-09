package internal

import (
	"context"

	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/Namchee/actions-case-police/internal/utils"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

// GithubClient is interface to interact with GitHub REST API v3
type GithubClient interface {
	GetIssue(context.Context, *entity.Meta, int) (*github.Issue, error)
	EditIssue(context.Context, *entity.Meta, int, *utils.IssueData) error
}

type githubClient struct {
	client *github.Client
}

// NewGithubClient creates a GitHub API client wrapper
func NewGithubClient(ctx context.Context, token string) GithubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	return &githubClient{client: github.NewClient(tc)}
}

func (cl *githubClient) GetIssue(
	ctx context.Context,
	meta *entity.Meta,
	number int,
) (*github.Issue, error) {
	issue, _, err := cl.client.Issues.Get(ctx, meta.Owner, meta.Name, number)

	return issue, err
}

func (cl *githubClient) EditIssue(
	ctx context.Context,
	meta *entity.Meta,
	number int,
	issue *utils.IssueData,
) error {
	_, _, err := cl.client.Issues.Edit(
		ctx,
		meta.Owner,
		meta.Name,
		number,
		&github.IssueRequest{
			Title: &issue.Title,
			Body:  &issue.Body,
		},
	)

	return err
}
