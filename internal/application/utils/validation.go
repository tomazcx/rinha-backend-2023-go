package utils

import (
	"errors"
	"time"

	"github.com/tomazcx/rinha-backend-go/internal/data/person"
)

var UnprocessableEntityError = errors.New("Invalid entry")

func ValidatePerson(dto person.CreatePessoaDTO) error {

	if len(dto.Name) > 100 || len(dto.Name) == 0 {	
		return UnprocessableEntityError
	}

	if len(dto.Nickname) > 32 || len(dto.Nickname) == 0 {	
		return UnprocessableEntityError
	}

	if _, err := time.Parse("2006-01-02", dto.Birthdate); err != nil {
		return UnprocessableEntityError
	}

	for _, tech := range dto.Stack {
		if len(tech) > 32 {
			return UnprocessableEntityError
		}
	}

	return nil
}
