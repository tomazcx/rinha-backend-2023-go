package peopledb

import (
	"database/sql"
	"errors"

	"github.com/tomazcx/rinha-backend-go/internal/entities"
)

type PersonRepository struct {
	db *sql.DB
}

func (r *PersonRepository) FindMany(t string) ([]entities.Person, error) {
	stmt, err := r.db.Prepare("SELECT id, nome, apelido, nascimento, stack FROM pessoa WHERE nome LIKE ? OR apelido LIKE ? OR stack LIKE ? LIMIT 50")
	if err != nil {
		return nil, err
	}

	term := "%" + t + "%"
	rows, err := stmt.Query(term, term, term)
	if err != nil {
		return nil, errors.New("Error preparing the statement")
	}

	var result []entities.Person
	for rows.Next() {
		var person entities.Person
		if err = rows.Scan(&person.ID, &person.Name, &person.Nickname, &person.Birthdate, &person.Stack); err != nil {
			return nil, err
		}
		result = append(result, person)
	}
	return result, nil
}

func (r *PersonRepository) FindById(id string) (*entities.Person, error) {
	stmt, err := r.db.Prepare("SELECT id, nome, apelido, nascimento, stack FROM pessoa WHERE id=?")
	if err != nil {
		return nil, errors.New("Error preparing the statement")
	}

	var person *entities.Person
	err = stmt.QueryRow().Scan(person.ID, person.Name, person.Nickname, person.Birthdate, person.Stack)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (r *PersonRepository) Create(person *entities.Person) error {
	stmt, err := r.db.Prepare("INSERT INTO pessoa (id, nome, apelido, nasicmento, stack) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("Error preparing the statement")
	}

	_, err = stmt.Exec(person.ID, person.Name, person.Nickname, person.Birthdate, person.StackStr())

	return err
}

func (r *PersonRepository) Count() (int, error) {
	stmt, _ := r.db.Prepare("SELECT COUNT(*) FROM pessoa")
	var count int
	err := stmt.QueryRow().Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PersonRepository) CheckNicknameTaken(nickname string) (bool, error) {
	stmt, err := r.db.Prepare("SELECT * FROM pessoa WHERE apelido = ?")

	if err != nil {
		return false, errors.New("Error preparing the statement")
	}

	_, err = stmt.Exec(nickname)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}
