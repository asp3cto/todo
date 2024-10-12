package api

type GitHubIssue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
}
