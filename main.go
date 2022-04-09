package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Namchee/actions-case-police/internal"
	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/Namchee/actions-case-police/internal/repository"
	"github.com/Namchee/actions-case-police/internal/utils"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lmsgprefix)
}

func main() {
	ctx := context.Background()

	cfg, err := entity.ReadConfiguration()
	if err != nil {
		logger.Fatalln(
			fmt.Sprintf("Failed to read action configuration: %w", err),
		)
	}

	cwd, _ := os.Getwd()

	meta, err := entity.CreateMeta(
		utils.ReadEnvString("GITHUB_REPOSITORY"),
	)
	if err != nil {
		logger.Fatalln(
			fmt.Sprintf("Failed to read metadata: %w", err),
		)
	}
	event, err := entity.ReadEvent(os.DirFS("/"))
	if err != nil {
		logger.Fatalln(
			fmt.Sprintf("Failed to read repository event: %w", err),
		)
	}

	client := internal.NewGithubClient(ctx, cfg.Token)

	issue, err := client.GetIssue(ctx, meta, event.Number)
	if err != nil {
		logger.Fatalln(
			fmt.Sprintf("Failed to get issue data: %w", err),
		)
	}

	dictionary := repository.GetDictionary(
		os.DirFS(fmt.Sprintf("%s/%s", cwd, "dict")),
		cfg.Preset,
	)

	if len(cfg.Dictionary) > 0 {
		utils.MergeDictionary(&cfg.Dictionary, &dictionary)
	}

	if len(cfg.Exclude) > 0 {
		utils.RemoveEntries(&cfg.Dictionary, cfg.Exclude)
	}

	result := utils.PolicizeIssue(issue, cfg.Dictionary)
}
