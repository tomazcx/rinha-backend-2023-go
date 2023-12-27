package peopledb

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/tomazcx/rinha-backend-go/internal/entities"
)

type PersonRepository struct {
	db *sql.DB
}

func (r *PersonRepository) FindMany(t string) ([]entities.Person, error) {
	stmt, err := r.db.Prepare("SELECT id, nome, apelido, nascimento, stack FROM pessoa WHERE searchable LIKE '%' || $1 || '%' LIMIT 50")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(t)
	if err != nil {
		return nil, errors.New("Error preparing the statement")
	}

	var result []entities.Person
	for rows.Next() {
		var person entities.Person
		var stackStr string
		if err = rows.Scan(&person.ID, &person.Name, &person.Nickname, &person.Birthdate, &stackStr); err != nil {
			return nil, err
		}
		person.Stack = strings.Split(stackStr, ",")
		result = append(result, person)
	}
	return result, nil
}

func (r *PersonRepository) FindById(id string) (*entities.Person, error) {
	stmt, err := r.db.Prepare("SELECT id, nome, apelido, nascimento, stack FROM pessoa WHERE id=$1")
	if err != nil {
		return nil, errors.New("Error preparing the statement")
	}

	var stackStr string
	var person entities.Person
	err = stmt.QueryRow(id).Scan(&person.ID, &person.Name, &person.Nickname, &person.Birthdate, &stackStr)
	if err != nil {
		return nil, err
	}

	person.Stack = strings.Split(stackStr, ",")

	return &person, nil
}

func (r *PersonRepository) Create(person *entities.Person) error {
	stmt, err := r.db.Prepare("INSERT INTO pessoa (id, nome, apelido, nascimento, stack) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return errors.New("Error preparing the statement")
	}

	_, err = stmt.Exec(person.ID, person.Name, person.Nickname, person.Birthdate, person.StackStr())

	return err
}

func (r *PersonRepository) Count() (int, error) {
	stmt, _ := r.db.Prepare("SELECT COUNT(1) FROM pessoa")
	var count int
	err := stmt.QueryRow().Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PersonRepository) CheckNicknameTaken(nickname string) (bool, error) {
	stmt, err := r.db.Prepare("SELECT COUNT(1) FROM pessoa WHERE apelido = $1")

	if err != nil {
		return false, errors.New("Error preparing the statement")
	}

	var nicknameTaken int
	_ = stmt.QueryRow(nickname).Scan(&nicknameTaken)

	if err != nil {
		return false, err
	}

	return nicknameTaken == 1, nil
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}
