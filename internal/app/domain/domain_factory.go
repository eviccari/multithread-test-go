package domain

import (
	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
	"github.com/eviccari/multithread-test-go/internal/app/domain/models"
)

func BuildUserModelFromDTO(dto dtos.GithubUserDTO) (user models.GithubUser) {
	return models.GithubUser{
		Login:  dto.Login,
		ID:     dto.ID,
		URL:    dto.URL,
		NodeID: dto.NodeID,
	}
}

func BuildDTOFromUserModel(model models.GithubUser) (dto dtos.GithubUserDTO) {
	return dtos.GithubUserDTO{
		Login:  model.Login,
		ID:     model.ID,
		URL:    model.URL,
		NodeID: model.NodeID,
	}
}
