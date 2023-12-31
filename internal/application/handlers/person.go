package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tomazcx/rinha-backend-go/internal/application/utils"
	"github.com/tomazcx/rinha-backend-go/internal/data/factory"
	"github.com/tomazcx/rinha-backend-go/internal/data/person"
)

type PersonHandler struct {
	factory factory.PersonFactory
}

func (h *PersonHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("t")

	if len(term) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	findManyPerson := h.factory.FindManyPeopleUseCase()
	person, err := findManyPerson.Execute(term)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := uuid.Parse(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	findPersonById := h.factory.FindPersonByIdUseCase()
	person, err := findPersonById.Execute(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto person.CreatePessoaDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := utils.ValidatePerson(dto); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	createPerson := h.factory.CreatePersonUseCase()
	createdPerson, err := createPerson.Execute(dto)

	if err != nil {
		if errors.Is(err, person.ErrNicknameAlreadyRegistered) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", "/pessoas/" + createdPerson.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPerson)
}

func (h *PersonHandler) GetCount(w http.ResponseWriter, r *http.Request) {
	countPerson := h.factory.CountPeopleUseCase()
	count, err := countPerson.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(count)))
}

func NewPersonHandler(factory factory.PersonFactory) *PersonHandler {
	return &PersonHandler{factory:factory}
} 
