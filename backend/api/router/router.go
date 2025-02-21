package router

import (
	"context"
    "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"
    "github.com/jackc/pgx/v5/pgxpool"

	"api-meli/api/resource/handlers"
)


func NewRouter(ctx context.Context, dbPool *pgxpool.Pool) *chi.Mux {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(common)

    // Basic CORS
    // for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
    router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"}, // allow anyone
        // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
        //AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // gotta review these before prod
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    }))

    router.Route("/v1", func(router chi.Router) { //root route
        //routes for datatype
        router.Route("/product", func(router chi.Router) {
            router.Get("/search", handlers.ReadByProductName(ctx, dbPool))
            router.Post("/", handlers.CreateProduct(ctx, dbPool))
            router.Get("/", handlers.ReadProduct(ctx, dbPool))
            router.Get("/{id}", handlers.ReadByProductId(ctx, dbPool))
            router.Put("/{id}", handlers.UpdateProduct(ctx, dbPool))
            router.Delete("/{id}", handlers.DeleteProduct(ctx, dbPool))
        })
        router.Route("/customer", func(router chi.Router) {
            router.Post("/", handlers.CreateCustomer(ctx, dbPool))
            router.Get("/", handlers.ReadCustomer(ctx, dbPool))
            router.Get("/{id}", handlers.ReadByCustomerId(ctx, dbPool))
            router.Put("/{id}", handlers.UpdateCustomer(ctx, dbPool))
            router.Delete("/{id}", handlers.DeleteCustomer(ctx, dbPool))
        })
        router.Route("/cart", func(router chi.Router) {
            router.Post("/", handlers.CreateCart(ctx, dbPool))
            router.Get("/", handlers.ReadCart(ctx, dbPool))
            router.Get("/{id}", handlers.ReadByCartId(ctx, dbPool))
            router.Put("/{id}", handlers.UpdateCart(ctx, dbPool))
            router.Delete("/{id}", handlers.DeleteCart(ctx, dbPool))
        })
        router.Route("/purchase", func(router chi.Router) {
            router.Post("/", handlers.CreatePurchase(ctx, dbPool))
            router.Get("/", handlers.ReadPurchase(ctx, dbPool))
            router.Get("/{id}", handlers.ReadByPurchaseId(ctx, dbPool))
            router.Put("/{id}", handlers.UpdatePurchase(ctx, dbPool))
            router.Delete("/{id}", handlers.DeletePurchase(ctx, dbPool))
        })
    })

    return router
}

// common configs and whategver
func common(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
