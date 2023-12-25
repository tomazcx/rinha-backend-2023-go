package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
		http.Error(w, "Invalid term", http.StatusBadRequest)
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	findPersonById := h.factory.FindPersonByIdUseCase()
	person, err := findPersonById.Execute(id)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto person.CreatePessoaDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid entry", http.StatusUnprocessableEntity)
		return
	}

	if err := utils.ValidatePerson(dto); err != nil {
		http.Error(w, "Invalid entry", http.StatusUnprocessableEntity)
		return
	}

	createPerson := h.factory.CreatePersonUseCase()
	createdPerson, err := createPerson.Execute(dto)

	if err != nil {
		if errors.Is(err, person.ErrNicknameAlreadyRegistered) {
			http.Error(w, "Nickname already registered", http.StatusConflict)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("/person/%s", createdPerson.ID))
	json.NewEncoder(w).Encode(createdPerson)
}

func (h *PersonHandler) GetCount(w http.ResponseWriter, r *http.Request) {
	countPerson := h.factory.CountPeopleUseCase()
	count, err := countPerson.Execute()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(count)))
}

func NewPersoHandler(factory factory.PersonFactory) *PersonHandler {
	return &PersonHandler{factory:factory}
}
