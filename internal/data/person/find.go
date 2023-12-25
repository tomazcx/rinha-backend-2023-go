package person

import (
	"github.com/tomazcx/rinha-backend-go/internal/data/protocol"
	"github.com/tomazcx/rinha-backend-go/internal/entities"
)

type FindPersonByIdUseCase struct {
	repo protocol.IPersonRepository
}

func (uc *FindPersonByIdUseCase) Execute(id string) (*entities.Person, error) {
	person, err := uc.repo.FindById(id)

	if err != nil {
		return nil, ErrPersonNotFound
	}

	return person, nil
}

func NewFindPersonByIdUseCase(repo protocol.IPersonRepository) *FindPersonByIdUseCase {
	return &FindPersonByIdUseCase{repo:repo}
}
