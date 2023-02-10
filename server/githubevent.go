package server

type GithubEvent struct {
	Action     string
	Ref        string           `json:"ref"`
	Before     string           `json:"before"`
	After      string           `json:"after"`
	Created    bool             `json:"created"`
	Deleted    bool             `json:"deleted"`
	Forced     bool             `json:"forced"`
	BaseRef    string           `json:"base_ref"`
	Compare    string           `json:"compare"`
	Commits    []GithubCommit   `json:"commits"`
	HeadCommit GithubCommit     `json:"head_commit"`
	Repository GithubRepository `json:"repository"`
	Pusher     GithubUser       `json:"pusher"`
	Sender     GithubUser       `json:"sender"`
}

type GithubCommit struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Timestamp string     `json:"timestamp"`
	URL       string     `json:"url"`
	Author    GithubUser `json:"author"`
	Committer GithubUser `json:"committer"`
}

type GithubRepository struct {
	ID       int        `json:"id"`
	NodeID   string     `json:"node_id"`
	Name     string     `json:"name"`
	FullName string     `json:"full_name"`
	Private  bool       `json:"private"`
	Owner    GithubUser `json:"owner"`
}

type GithubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
