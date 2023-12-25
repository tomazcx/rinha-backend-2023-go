package application

import (
	"github.com/go-chi/chi"
	"github.com/tomazcx/rinha-backend-go/internal/application/handlers"
	"github.com/tomazcx/rinha-backend-go/internal/data/factory"
)

type ApplicationRouter struct {
	handler *handlers.PersonHandler
}

func (a *ApplicationRouter) DefineRoutes(r *chi.Mux) {
	r.Post("/pessoas", a.handler.Create)
	r.Get("/pessoas/{id}", a.handler.GetOne)
	r.Get("/pessoas", a.handler.GetMany)
	r.Get("/contagem-pessoas", a.handler.GetCount)
}

func NewApplicationRouter() *ApplicationRouter {
	factory := factory.PersonFactory{}
	handler := handlers.NewPersonHandler(factory)
	return &ApplicationRouter{
		handler: handler,
	}
}
