package service

import (
	"github.com/Namchee/actions-case-police/internal/entity"
	"github.com/fatih/color"
)

// LogResult logs the result to console
func LogResult(
	result *entity.IssueData,
	cfg *entity.Configuration,
) {
	green := color.New(color.FgCyan)

	green.Println("🚨 case-police executed successfully")

	if len(result.Changes) > 0 {
		if cfg.Fix {
			green.Println("Below are all possible fixes by case-police")
		} else {
			green.Println("Below are all changes introduces by case-police")
		}

		yellow := color.New(color.FgYellow)

		for k, v := range result.Changes {
			yellow.Printf("%s ➔", k)
			green.Printf("%s\n", v)
		}
	} else {
		green.Println("All cases are already correct :)")
	}
}
