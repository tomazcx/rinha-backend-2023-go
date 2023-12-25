package person

import (
	"github.com/tomazcx/rinha-backend-go/internal/data/protocol"
	"github.com/tomazcx/rinha-backend-go/internal/entities"
)

type CreatePessoaDTO struct {
	Name      string   `json:"nome"`
	Nickname  string   `json:"apelido"`
	Birthdate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

type CreatePerson struct {
	repo protocol.IPersonRepository
}

func (uc *CreatePerson) Execute(dto CreatePessoaDTO) (*entities.Person, error) {

	nicknameTaken, err := uc.repo.CheckNicknameTaken(dto.Nickname)

	if err != nil {
		return nil, err
	}

	if nicknameTaken {
		return nil, ErrNicknameAlreadyRegistered
	}

	person := entities.NewPerson(dto.Name, dto.Nickname, dto.Stack, dto.Birthdate)
	err = uc.repo.Create(person)

	if err != nil {
		return nil, err
	}

	return person, nil
}

func NewCreatePerson(repo protocol.IPersonRepository) *CreatePerson {
	return &CreatePerson{repo: repo}
}
