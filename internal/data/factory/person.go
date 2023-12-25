package factory

import (
	"github.com/tomazcx/rinha-backend-go/internal/data/person"
	"github.com/tomazcx/rinha-backend-go/internal/infra/db"
	"github.com/tomazcx/rinha-backend-go/internal/infra/db/peopledb"
)

type PersonFactory struct{}

func (f PersonFactory) FindPersonByIdUseCase() *person.FindPersonByIdUseCase{
	db := db.GetDBConn()
	repo := peopledb.NewPersonRepository(db)
	return person.NewFindPersonByIdUseCase(repo)
}

func (f PersonFactory) FindManyPeopleUseCase() *person.FindManyPeopleUseCase{
	db := db.GetDBConn()
	repo := peopledb.NewPersonRepository(db)
	return person.NewFindManyPeopleUseCase(repo)
}

func (f PersonFactory) CreatePersonUseCase() *person.CreatePerson {
	db := db.GetDBConn()
	repo := peopledb.NewPersonRepository(db)
	return person.NewCreatePerson(repo)
}

func (f PersonFactory) CountPeopleUseCase() *person.CountPersonUseCase {
	db := db.GetDBConn()
	repo := peopledb.NewPersonRepository(db)
	return person.NewCountPersonUseCase(repo)
}
