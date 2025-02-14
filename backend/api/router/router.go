package router

import (
    "github.com/go-chi/chi/v5"
    "api-meli/api/resource/datatype"
)


func NewRouter() *chi.Mux {
    router := chi.NewRouter()

    router.Route("/v1", func(router chi.Router) { //root route
        //routes for datatype
        router.Route("/datatype/", func(router chi.Router) {
            router.Post("/", datatype.Create)
            router.Get("/", datatype.Read)
            router.Get("/{id}", datatype.ReadById)
            router.Put("/{id}", datatype.UpdateById)
            router.Delete("/{id}", datatype.DeleteById)
        })
    })

    return router
}
