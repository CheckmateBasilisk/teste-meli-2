//	this package contains the handlers for the type 'customer'
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
//  @summary        Create Customer
//  @description    Create Customer
//  @tags           Customer
//  @accept         json
//  @produce        json
//  @param          body    body    repositories.Customer    true    "Customer form"
//  @success        201
//  @failure        400 {object}    errors.Error
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Customer [post]
func CreateCustomer(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        //getting data
        var x repositories.CreateCustomerParams
        err := json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //running
        p, err := repo.CreateCustomer(ctx, x)
        if err != nil {
            msg, code := errors.DbInsertFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //writing results to w
        err = json.NewEncoder(w).Encode(p)
        if err != nil {
            msg, code := errors.JsonEncodeFailure() //FIXME: here problem is server-side not client-side...
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

// Read godoc
//
//  @summary        Read all Customer
//  @description    Read all Customer
//  @tags           Customer
//  @accept         json
//  @produce        json
//  @success        200 {array}     repositories.Customer
//  @failure        500 {object}    errors.Error
//  @router         /Customer [get]
func ReadCustomer(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        ps, err := repo.ReadCustomer(ctx)
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
//  @summary        Read Customer by id
//  @description    Read Customer by id
//  @tags           Customer
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "Customer ID"
//  @success        200 {object}    repositories.Customer
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Customer/{id} [get]
func ReadByCustomerId(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
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
        p, err := repo.ReadByCustomerId(ctx, uuid)
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
//  @summary        Update Customer
//  @description    Update Customer
//  @tags           Customer
//  @accept         json
//  @produce        json
//  @param          id      path    string  true    "Customer ID"
//  @param          body    body    repositories.Customer    true    "Customer form"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Customer/{id} [put]
func UpdateCustomer(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
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

        var x repositories.UpdateCustomerParams
        err = json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        x.ID = uuid
        p , err := repo.UpdateCustomer(ctx, x)
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
//  @summary        Delete Customer
//  @description    Delete Customer
//  @tags           Customer
//  @accept         json
//  @produce        json
//  @param          id  path    string  true    "Customer ID"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Customer/{id} [delete]
func DeleteCustomer(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
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

        //call db.DeleteCustomer(ctx context.Context, conn, arg)
        err = repo.DeleteCustomer(ctx, uuid)
        if err != nil {
            msg, code := errors.DbDeleteFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

