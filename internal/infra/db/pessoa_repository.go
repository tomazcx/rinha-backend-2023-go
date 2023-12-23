package repository

import (
	"database/sql"
	"errors"

	"github.com/tomazcx/rinha-backend-go/internal/entities"
)

type PessoaRepository struct {
	db *sql.DB
}

func NewPessoaRepository(db *sql.DB) *PessoaRepository{
	return &PessoaRepository{db: db}
}

func (r *PessoaRepository) FindMany(name, apelido, stack string) ([]entities.Pessoa, error) {
	stmt, err := r.db.Prepare("SELECT id, nome, apelido, nascimento, stack FROM pessoa WHERE nome LIKE ? OR apelido LIKE ? OR stack LIKE ?")
	if err != nil {
		return nil, err
	}

	likeName := "%" + name + "%"
	likeApelido := "%" + apelido + "%"
	likeStack := "%" + stack + "%"
	rows, err := stmt.Query(likeName, likeApelido, likeStack)
	if err != nil {
		return nil, errors.New("Error preparing the statement")
	}

	var result []entities.Pessoa
	for rows.Next() {
		var pessoa entities.Pessoa
		if err = rows.Scan(&pessoa.ID, &pessoa.Nome, &pessoa.Apelido, &pessoa.Nascimento, &pessoa.Stack); err != nil {
			return nil, err
		}
		result = append(result, pessoa)
	}
	return result, nil
}

func (r *PessoaRepository) FindById(id string) (*entities.Pessoa, error) {
	stmt, err := r.db.Prepare("SELECT id, nome, apelido, nascimento, stack FROM pessoa WHERE id=?")
	if err != nil {
		return nil, errors.New("Error preparing the statement")
	}

	var pessoa *entities.Pessoa
	err = stmt.QueryRow().Scan(pessoa.ID, pessoa.Nome, pessoa.Apelido, pessoa.Nascimento, pessoa.Stack)
	if err != nil {
		return nil, err
	}

	return pessoa, nil
}

func (r *PessoaRepository) Create(pessoa *entities.Pessoa) error {
	stmt, err := r.db.Prepare("INSERT INTO pessoa (id, nome, apelido, nasicmento, stack) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("Error preparing the statement")
	}

	_, err = stmt.Exec(pessoa.ID, pessoa.Nome, pessoa.Apelido, pessoa.Nascimento, pessoa.Stack)

	return err	
}

func (r *PessoaRepository) Count() (int, error) {
	stmt, _ := r.db.Prepare("SELECT COUNT(*) FROM pessoa")
	var count int	
	err := stmt.QueryRow().Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}


