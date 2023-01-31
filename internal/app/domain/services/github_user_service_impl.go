package services

import (
	"log"

	"github.com/eviccari/multithread-test-go/internal/adapters"
	"github.com/eviccari/multithread-test-go/internal/app/domain"
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
	dirtyList, err := gusi.r.Get(pageSize, since)
	if err != nil {
		return
	}

	for _, dto := range dirtyList {
		model := domain.BuildGithubUserModelFromDTO(dto)
		if el := model.Validate(); len(el) == 0 {
			githubUsers = append(githubUsers, dto)
		} else {
			log.Printf("\n errors on load github login: %s -> %v", dto.Login, el)
		}
	}

	return
}
