//	this package contains the handlers for the type 'cart'
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
	"github.com/jackc/pgx/v5/pgxpool"

	"api-meli/api/resource/repositories"
	"api-meli/api/common/errors"
)

// Create godoc
//
//  @summary        Create Cart
//  @description    Create Cart
//  @tags           Cart
//  @accept         json
//  @produce        json
//  @param          body    body    repositories.Cart    true    "Cart form"
//  @success        201
//  @failure        400 {object}    errors.Error
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Cart [post]
func CreateCart(ctx context.Context, conn *pgxpool.Pool) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        //getting data
        var x repositories.CreateCartParams
        err := json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //running
        c, err := repo.CreateCart(ctx, x)
        if err != nil {
            msg, code := errors.DbInsertFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //writing results to w
        err = json.NewEncoder(w).Encode(c)
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
//  @summary        Read all Cart
//  @description    Read all Cart
//  @tags           Cart
//  @accept         json
//  @produce        json
//  @success        200 {array}     repositories.Cart
//  @failure        500 {object}    errors.Error
//  @router         /Cart [get]
func ReadCart(ctx context.Context, conn *pgxpool.Pool) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        cs, err := repo.ReadCart(ctx)
        if err != nil {
            msg, code := errors.DbReadFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        err = json.NewEncoder(w).Encode(cs)
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
//  @summary        Read Cart by id
//  @description    Read Cart by id
//  @tags           Cart
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "Cart ID"
//  @success        200 {object}    repositories.Cart
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Cart/{id} [get]
func ReadByCartId(ctx context.Context, conn *pgxpool.Pool) http.HandlerFunc {
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
        p, err := repo.ReadByCartId(ctx, uuid)
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
//  @summary        Update Cart
//  @description    Update Cart
//  @tags           Cart
//  @accept         json
//  @produce        json
//  @param          id      path    string  true    "Cart ID"
//  @param          body    body    repositories.Cart    true    "Cart form"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Cart/{id} [put]
func UpdateCart(ctx context.Context, conn *pgxpool.Pool) http.HandlerFunc {
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

        var x repositories.UpdateCartParams
        err = json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        x.ID = uuid
        c , err := repo.UpdateCart(ctx, x)
        if err != nil {
            msg, code := errors.DbUpdateFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        err  = json.NewEncoder(w).Encode(c)
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
//  @summary        Delete Cart
//  @description    Delete Cart
//  @tags           Cart
//  @accept         json
//  @produce        json
//  @param          id  path    string  true    "Cart ID"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Cart/{id} [delete]
func DeleteCart(ctx context.Context, conn *pgxpool.Pool) http.HandlerFunc {
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

        //call db.DeleteCart(ctx context.Context, conn, arg)
        err = repo.DeleteCart(ctx, uuid)
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
//  @summary        Read Cart by customer_id
//  @description    Read Cart by customer_id
//  @tags           Cart
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "ID"
//  @success        200 {object}    repositories.Cart
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /customer/{id}/cart/ [get]
func ReadCartByCustomerId(ctx context.Context, conn *pgxpool.Pool) http.HandlerFunc {
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
        p, err := repo.ReadCartByCustomerId(ctx, uuid)
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

