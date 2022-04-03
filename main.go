package main

import (
	"github.com/google/go-github/v43/github"
)

var (
	dictionaries = []string{
		"abbreviates.json",
		"brands.json",
		"general.json",
		"products.json",
		"softwares.json",
	}
)

func main() {
	client := github.NewClient()
}
