package repositories

import (
	"errors"
	"fmt"

	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

type GithubUserMockRepository struct{}

func NewGithubUserMockRepository() GithubUserMockRepository {
	return GithubUserMockRepository{}
}

func (gur GithubUserMockRepository) Get(pageSize, since int) (githubUsers []dtos.GithubUserDTO, err error) {
	if pageSize < 1 || since < 0 {
		err = errors.New("pageSize must be greater than 0.\n since must be greater than or equal 0")
		return
	}

	next := since + 1

	for i := 0; i < pageSize; i++ {
		githubLogin := fmt.Sprintf("GithubLogin_%d", next)

		githubUsers = append(githubUsers, dtos.GithubUserDTO{
			ID:      next,
			NodeID:  fmt.Sprintf("NodeID_%d", next),
			Login:   githubLogin,
			URL:     fmt.Sprintf("https://api.github.com/%s/api", githubLogin),
			HTMLURL: fmt.Sprintf("https://github.com/%s", githubLogin),
		})

		next++
	}

	return
}
