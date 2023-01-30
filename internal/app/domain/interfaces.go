package domain

import "github.com/eviccari/multithread-test-go/internal/app/domain/dtos"

type UserService interface {
	Load() (usersList []dtos.GithubUserDTO, err error)
	Transform(userData dtos.GithubUserDTO) (updatedUser dtos.GithubUserDTO, err error)
	GenerateOutput(userData dtos.GithubUserDTO) (updatedUser dtos.GithubUserDTO, err error)
}

type GithubUserService interface {
	Get(pageSize, since int) (githubUsers []dtos.GithubUserDTO, err error)
}

type Orchestrator interface {
	Execute() (errorsList []error)
}
