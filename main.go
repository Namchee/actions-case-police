package main

import (
	"fmt"
	"os"

	"github.com/Namchee/actions-case-police/internal"
	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/Namchee/actions-case-police/internal/repository"
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
	cfg, err := entity.ReadConfiguration()
	if err != nil {
		errorLogger.Println(err)
		os.Exit(1)
	}

	cwd, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf("%s/%s", cwd, "dict"))

	dictionary := repository.GetDictionary(fsys, cfg.Preset)

	client := internal.NewGithubClient(cfg.Token)
}
