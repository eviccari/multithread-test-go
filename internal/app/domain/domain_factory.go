package domain

import (
	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
	"github.com/eviccari/multithread-test-go/internal/app/domain/models"
)

func BuildGithubUserModelFromDTO(dto dtos.GithubUserDTO) (user models.GithubUser) {
	return models.GithubUser{
		Login:  dto.Login,
		ID:     dto.ID,
		URL:    dto.URL,
		NodeID: dto.NodeID,
	}
}

func BuildGithubUserDTOFromModel(model models.GithubUser) (dto dtos.GithubUserDTO) {
	return dtos.GithubUserDTO{
		Login:  model.Login,
		ID:     model.ID,
		URL:    model.URL,
		NodeID: model.NodeID,
	}
}

func BuildGreatestUserModelFromDTO(dto dtos.GreatestUserDTO) (user models.GreatestUser) {
	return models.GreatestUser{
		ID:            dto.ID,
		LegacyID:      dto.LegacyID,
		LegacyLogin:   dto.LegacyLogin,
		LegacyNodeID:  dto.LegacyNodeID,
		LegacyURL:     dto.LegacyURL,
		LegacyHTMLURL: dto.LegacyHTMLURL,
		NewEmail:      dto.NewEmail,
		CreatedAt:     dto.CreatedAt,
	}
}

func BuildGreatestUserDTOFromModel(model models.GreatestUser) (dto dtos.GreatestUserDTO) {
	return dtos.GreatestUserDTO{
		ID:            model.ID,
		LegacyID:      model.LegacyID,
		LegacyLogin:   model.LegacyLogin,
		LegacyNodeID:  model.LegacyNodeID,
		LegacyURL:     model.LegacyURL,
		LegacyHTMLURL: model.LegacyHTMLURL,
		NewEmail:      model.NewEmail,
		CreatedAt:     model.CreatedAt,
	}
}

func BuildGreatestUserDTOFromGithubUserDTO(guDTO dtos.GithubUserDTO) (dto dtos.GreatestUserDTO) {
	return dtos.GreatestUserDTO{
		LegacyLogin:   guDTO.Login,
		LegacyID:      guDTO.ID,
		LegacyURL:     guDTO.URL,
		LegacyHTMLURL: guDTO.HTMLURL,
	}
}
