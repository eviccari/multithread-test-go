package models

type GithubUser struct {
	Login   string
	ID      int
	NodeID  string
	URL     string
	HTMLURL string
}

func NewGithubUser(login string, id int, nodeID, url, HTMLURL string) (user GithubUser) {
	return GithubUser{
		Login:  login,
		ID:     id,
		NodeID: nodeID,
		URL:    url,
	}
}

func (gu GithubUser) Validate() (errorsList []error) {
	pm := NewGithubUserPoliciesManager(gu)
	return pm.Apply()
}
