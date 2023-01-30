package adapters

import "github.com/eviccari/multithread-test-go/internal/app/domain/dtos"

type GithubUserRepository interface {
	Get(pageSize, since int) (githubUsers []dtos.GithubUserDTO, err error)
}
