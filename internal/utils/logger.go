package utils

import "github.com/fatih/color"

// LogResult logs the result to console
func LogResult(
	result *IssueData,
) {
	green := color.New(color.FgCyan)

	green.Println("🚨 case-police executed successfully")

	if len(result.Changes) > 0 {
		yellow := color.New(color.FgYellow)

		for k, v := range result.Changes {
			yellow.Printf("%s ➔", k)
			green.Printf("%s\n", v)
		}
	}

	green.Println("All cases are correct now :)")
}
