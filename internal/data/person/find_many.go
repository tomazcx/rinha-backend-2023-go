package person

import (
	"github.com/tomazcx/rinha-backend-go/internal/data/protocol"
	"github.com/tomazcx/rinha-backend-go/internal/entities"
)

type FindManyPeopleUseCase struct {
	repo protocol.IPersonRepository
}

func (uc *FindManyPeopleUseCase) Execute(t string) ([]entities.Person, error) {
	people, err := uc.repo.FindMany(t)

	if err != nil {
		return nil, err
	}

	return people, nil
}

func NewFindManyPeopleUseCase(repo protocol.IPersonRepository) *FindManyPeopleUseCase {
	return &FindManyPeopleUseCase{repo:repo}
}
