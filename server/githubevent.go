package server

type GitHubEvent struct {
	Action     string `json:"action"`
	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		URL      string `json:"html_url"`
	} `json:"repository"`
	Sender struct {
		Login string `json:"login"`
		URL   string `json:"html_url"`
	} `json:"sender"`
}
