package repositories

import (
	"log"

	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

type GreatestUserMockRepository struct{}

func NewGreatestUserMockRepository() GreatestUserMockRepository {
	return GreatestUserMockRepository{}
}

func (gump GreatestUserMockRepository) Create(dto dtos.GreatestUserDTO) (err error) {
	log.Printf("user ID %s with legacy login %s created", dto.ID, dto.LegacyLogin)
	return
}

func (gump GreatestUserMockRepository) SetEmpty() (err error) {
	return
}
