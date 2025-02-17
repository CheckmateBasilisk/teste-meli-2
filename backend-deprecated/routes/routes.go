package routes

import (
	"net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

	"github.com/CheckmateBasilisk/teste-meli-2/backend/routes/handlers"
)

func BuildRouter() *chi.Mux {
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte(string("Hello, World!\n")))
    })

    //creates a subrouter for a given data type
    router.Route("/datatype", func(router chi.Router) {
        router.Post("/", handlers.CreateDataType)
        router.Get("/", handlers.ReadDataType)
        router.Get("/{id}", handlers.ReadDataTypeById)
        router.Put("/{id}", handlers.UpdateDataTypeById)
        router.Delete("/{id}", handlers.DeleteDataTypeById)
    })


    return router
}





