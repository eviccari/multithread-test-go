package dtos

type GithubUserDTO struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}
