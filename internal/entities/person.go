package entities

import (
	"strings"

	"github.com/google/uuid"
)

type Person struct {
	ID        string   `json:"id"`
	Name      string   `json:"nome"`
	Nickname  string   `json:"apelido"`
	Birthdate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

func (p *Person) StackStr() string {
	return strings.Join(p.Stack, ",")
}

func NewPerson(name string, nickname string, stack []string, birthdate string) *Person {
	uuid := uuid.New()
	return &Person{
		ID:        uuid.String(),
		Name:      name,
		Nickname:  nickname,
		Stack:     stack,
		Birthdate: birthdate,
	}
}
