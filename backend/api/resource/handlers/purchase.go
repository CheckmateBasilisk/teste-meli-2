//	this package contains the handlers for the type 'purchase'
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
//  @summary        Create Purchase
//  @description    Create Purchase
//  @tags           Purchase
//  @accept         json
//  @produce        json
//  @param          body    body    repositories.Purchase    true    "Purchase form"
//  @success        201
//  @failure        400 {object}    errors.Error
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Purchase [post]
func CreatePurchase(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        //getting data
        var x repositories.CreatePurchaseParams
        err := json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
        //running
        p, err := repo.CreatePurchase(ctx, x)
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
//  @summary        Read all Purchase
//  @description    Read all Purchase
//  @tags           Purchase
//  @accept         json
//  @produce        json
//  @success        200 {array}     repositories.Purchase
//  @failure        500 {object}    errors.Error
//  @router         /Purchase [get]
func ReadPurchase(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
    repo := repositories.New(conn)
    return func (w http.ResponseWriter, r *http.Request) {
        ps, err := repo.ReadPurchase(ctx)
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
//  @summary        Read Purchase by id
//  @description    Read Purchase by id
//  @tags           Purchase
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "Purchase ID"
//  @success        200 {object}    repositories.Purchase
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Purchase/{id} [get]
func ReadByPurchaseId(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
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
        p, err := repo.ReadByPurchaseId(ctx, uuid)
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
//  @summary        Update Purchase
//  @description    Update Purchase
//  @tags           Purchase
//  @accept         json
//  @produce        json
//  @param          id      path    string  true    "Purchase ID"
//  @param          body    body    repositories.Purchase    true    "Purchase form"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        422 {object}    errors.Error
//  @failure        500 {object}    errors.Error
//  @router         /Purchase/{id} [put]
func UpdatePurchase(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
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

        var x repositories.UpdatePurchaseParams
        err = json.NewDecoder(r.Body).Decode(&x)
        if err != nil {
            msg, code := errors.JsonDecodeFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }

        x.ID = uuid
        p , err := repo.UpdatePurchase(ctx, x)
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
//  @summary        Delete Purchase
//  @description    Delete Purchase
//  @tags           Purchase
//  @accept         json
//  @produce        json
//  @param          id  path    string  true    "Purchase ID"
//  @success        200
//  @failure        400 {object}    errors.Error
//  @failure        404
//  @failure        500 {object}    errors.Error
//  @router         /Purchase/{id} [delete]
func DeletePurchase(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
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

        //call db.DeletePurchase(ctx context.Context, conn, arg)
        err = repo.DeletePurchase(ctx, uuid)
        if err != nil {
            msg, code := errors.DbDeleteFailure()
            log.Println(fmt.Errorf(msg, err))
            http.Error(w, msg, code)
            return
        }
    }
}

