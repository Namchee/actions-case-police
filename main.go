package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Namchee/actions-case-police/internal"
	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/Namchee/actions-case-police/internal/repository"
	"github.com/Namchee/actions-case-police/internal/service"
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
			fmt.Errorf("Failed to read action configuration: %w", err),
		)
	}

	meta, err := entity.CreateMeta(
		utils.ReadEnvString("GITHUB_REPOSITORY"),
	)
	if err != nil {
		logger.Fatalln(
			fmt.Errorf("Failed to read metadata: %w", err),
		)
	}
	event, err := entity.ReadEvent(os.DirFS("/"))
	if err != nil {
		logger.Fatalln(
			fmt.Errorf("Failed to read repository event: %w", err),
		)
	}

	client := internal.NewGithubClient(ctx, cfg.Token)

	issue, err := client.GetIssue(ctx, meta, event.Number)
	if err != nil {
		logger.Fatalln(
			fmt.Errorf("Failed to get issue data: %w", err),
		)
	}

	actionPath := utils.ReadEnvString("GITHUB_ACTION_PATH")
	dictionary := repository.GetDictionary(
		os.DirFS(fmt.Sprintf("%s/%s", actionPath, "dict")),
		cfg.Preset,
	)

	utils.MergeDictionary(&cfg.Dictionary, &dictionary)
	if len(cfg.Exclude) > 0 {
		utils.RemoveEntries(&cfg.Dictionary, cfg.Exclude)
	}

	result := service.PolicizeIssue(issue, cfg)

	if len(result.Changes) > 0 {
		err = client.EditIssue(ctx, meta, event.Number, result)

		if err != nil {
			log.Fatalln(
				fmt.Errorf("Failed to edit issue: %w", err),
			)
		}
	}

	service.LogResult(result, cfg)
}
