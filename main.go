package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Namchee/actions-case-police/internal"
	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/Namchee/actions-case-police/internal/repository"
	"github.com/Namchee/actions-case-police/internal/utils"
	"github.com/fatih/color"
)

var (
	infoLogger  color.Color
	errorLogger color.Color
)

func init() {
	infoLogger = *color.New(color.FgYellow)
	errorLogger = *color.New(color.FgRed)
}

func main() {
	ctx := context.Background()

	cfg, err := entity.ReadConfiguration()
	if err != nil {
		errorLogger.Println(err)
		os.Exit(1)
	}

	cwd, _ := os.Getwd()

	meta, err := entity.CreateMeta(
		utils.ReadEnvString("GITHUB_REPOSITORY"),
	)
	if err != nil {
		errorLogger.Printf("Failed to read repository metadata: %s", err.Error())
		os.Exit(1)
	}
	event, err := entity.ReadEvent(os.DirFS("/"))
	if err != nil {
		errorLogger.Printf("Failed to read repository event: %s", err.Error())
		os.Exit(1)
	}

	client := internal.NewGithubClient(ctx, cfg.Token)

	issue, err := client.GetIssue(ctx, meta, event.Number)
	if err != nil {
		errorLogger.Printf("Failed to get issue: %s", err.Error())
		os.Exit(1)
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
}
