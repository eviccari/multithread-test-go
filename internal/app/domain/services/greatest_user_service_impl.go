package services

import (
	"github.com/eviccari/multithread-test-go/internal/adapters"
	"github.com/eviccari/multithread-test-go/internal/app/domain"
	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

type GreatestUserServiceImpl struct {
	repo adapters.GreatestUserRepository
}

func NewGreatestUserServiceImpl(repo adapters.GreatestUserRepository) GreatestUserServiceImpl {
	return GreatestUserServiceImpl{
		repo: repo,
	}
}

func (gusi GreatestUserServiceImpl) Create(dto dtos.GreatestUserDTO) (errorsList []error) {
	model := domain.BuildGreatestUserModelFromDTO(dto)
	model.SetAsNew()
	errorsList = model.Validate()

	if len(errorsList) > 0 {
		return
	}

	dto = domain.BuildGreatestUserDTOFromModel(model)

	if err := gusi.repo.Create(dto); err != nil {
		errorsList = append(errorsList, err)
	}

	return
}

func (gusi GreatestUserServiceImpl) SetEmpty() (err error) {
	return gusi.repo.SetEmpty()
}
