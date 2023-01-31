package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

type GithubUserHTTPRepository struct{}

func NewGitHubUserHTTPRepository() GithubUserHTTPRepository {
	return GithubUserHTTPRepository{}
}

func (gur GithubUserHTTPRepository) Get(pageSize, since int) (githubUsers []dtos.GithubUserDTO, err error) {
	if pageSize < 1 || since < 0 {
		err = errors.New("pageSize must be greater than 0.\n since must be greater than or equal 0")
		return
	}

	time.Sleep(15 * time.Second)

	githubURL := "https://api.github.com/users"
	res, err := http.Get(fmt.Sprintf("%s?per_page=%d&since=%d", githubURL, pageSize, since))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		err = json.NewDecoder(res.Body).Decode(&githubUsers)
	} else {
		err = errors.New(res.Status)
	}

	return
}
