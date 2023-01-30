package services

import (
	"github.com/eviccari/multithread-test-go/internal/adapters"
	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

type GithubUserServiceImpl struct {
	r adapters.GithubUserRepository
}

func NewGithubUserServiceImpl(r adapters.GithubUserRepository) GithubUserServiceImpl {
	return GithubUserServiceImpl{
		r: r,
	}
}

func (gusi GithubUserServiceImpl) Get(pageSize, since int) (githubUsers []dtos.GithubUserDTO, err error) {
	return gusi.r.Get(pageSize, since)
}
