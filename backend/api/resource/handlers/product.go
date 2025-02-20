//	this package contains the handlers for the type 'product'
//	the handlers are called @api/router/router.go
package handlers

import (
    "fmt"
    "log"
	"context"
	"net/http"
    "encoding/json"

    "github.com/google/uuid"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"api-meli/api/resource/repositories"
	"api-meli/api/common/errors"
)

// Create godoc
//
//  @summary        Create Product
//  @description    Create Product
//  @tags           Product
//  @accept         json
//  @produce        json
//  @param          body    body    repositories.Product    true    "Product form"
//  @success        201
//  @failure        400 {object}    errors.Error
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Product [post]
func CreateProduct(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        //getting data
        var x repositories.CreateProductParams
        err := json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //running
        p, err := repo.CreateProduct(ctx, x)
        if err != nil {
            msg, code := errors.DbInsertFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //writing results to w
        err = json.NewEncoder(w).Encode(p)
        if err != nil {
            msg, code := errors.JsonEncodeFailure() //FIXME: here problem is server-side not user-side...
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

// Read godoc
//
//  @summary        Read all Product
//  @description    Read all Product
//  @tags           Product
//  @accept         json
//  @produce        json
//  @success        200 {array}     repositories.Product
//  @failure        500 {object}    errors.Error
//  @router         /Product [get]
func ReadProduct(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        ps, err := repo.ReadProduct(ctx)
        if err != nil {
            msg, code := errors.DbReadFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        err = json.NewEncoder(w).Encode(ps)
         if err != nil {
            msg, code := errors.JsonEncodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

// Read godoc
//
//  @summary        Read Product by id
//  @description    Read Product by id
//  @tags           Product
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "Product ID"
//  @success        200 {object}    repositories.Product
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Product/{id} [get]
func ReadByProductId(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        const paramName = "id"
        id := chi.URLParam(r, paramName)
        uuid, err := uuid.Parse(id)
        if err != nil {
            msg, code := errors.InvalidURLParam(paramName)
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        p, err := repo.ReadByProductId(ctx, uuid)
        if err != nil {
            msg, code := errors.DbReadFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        err = json.NewEncoder(w).Encode(p)
        if err != nil {
            msg, code := errors.JsonEncodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

// Update godoc
//
//  @summary        Update Product
//  @description    Update Product
//  @tags           Product
//  @accept         json
//  @produce        json
//  @param          id      path    string  true    "Product ID"
//  @param          body    body    repositories.Product    true    "Product form"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Product/{id} [put]
func UpdateProduct(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        const paramName = "id"
        id := chi.URLParam(r, paramName)
        uuid, err := uuid.Parse(id)
        if err != nil {
            msg, code := errors.InvalidURLParam(paramName)
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        var x repositories.UpdateProductParams
        err = json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        x.ID = uuid
        p , err := repo.UpdateProduct(ctx, x)
        if err != nil {
            msg, code := errors.DbUpdateFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        err  = json.NewEncoder(w).Encode(p)
        if err != nil {
            msg, code := errors.JsonEncodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

// Delete godoc
//
//  @summary        Delete Product
//  @description    Delete Product
//  @tags           Product
//  @accept         json
//  @produce        json
//  @param          id  path    string  true    "Product ID"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Product/{id} [delete]
func DeleteProduct(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        const paramName = "id"
        id := chi.URLParam(r, paramName)
        uuid, err := uuid.Parse(id)
        if err != nil {
            msg, code := errors.InvalidURLParam(paramName)
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        //call db.DeleteProduct(ctx context.Context, conn, arg)
        err = repo.DeleteProduct(ctx, uuid)
        if err != nil {
            msg, code := errors.DbDeleteFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

// Read godoc
//
//  @summary        Read Product by name
//  @description    Read Product by name
//  @tags           Product
//  @accept         json
//  @produce        json
//  @param          name	query        string  true    "Product Name"
//  @success        200 {object}    repositories.Product
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /product/search     [get]
func ReadByProductName(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        const paramName = "name"

        name := r.URL.Query().Get(paramName)


        p, err := repo.ReadByProductName(ctx, name)
        if err != nil {
            msg, code := errors.DbReadFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        fmt.Println(">>>> ", name)
        fmt.Println(">>>> ", p)

        err = json.NewEncoder(w).Encode(p)
        if err != nil {
            msg, code := errors.JsonEncodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}
