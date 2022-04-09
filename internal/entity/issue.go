package entity

// IssueData represents fixed case GitHub issue
type IssueData struct {
	Title   string
	Body    string
	Changes map[string]string
}
