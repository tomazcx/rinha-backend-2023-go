package protocol

import "github.com/tomazcx/rinha-backend-go/internal/entities"

type IPersonRepository interface {
	Count() (int, error)
	Create(person *entities.Person) error
	FindById(id string) (*entities.Person, error)
	FindMany(t string) ([]entities.Person, error)
	CheckNicknameTaken(nickname string) (bool, error)
}
