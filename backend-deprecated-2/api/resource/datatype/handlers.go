// FIXME: THIS IS JUST A TEMPLATE
//      this package contains the handlers for the type 'datatype'
//      the handlers are called @api/router/router.go
package datatype

import (
    "net/http"
    "api-meli/api/common/err"
)

//FIXME: i dont understand why use this kind of type... seems weird af
// apparently I'll attach loggers, validators and Repository to this, giving handler.go (here) access to db etc
//    now, why cant these be used here directly im not exactly sure... manye it is an injection thing (seems to be so)?
//   anyways, need to add func (a *API) to the definition of each handler below later...
// a API instance has additional info for the handlers to work, its like passing it as argument, but this way it can be called as a method of an instace of API created in the scope above... or something like that
type API struct {}

// Create godoc
//
//  @summary        Create DataType
//  @description    Create DataType
//  @tags           DataType
//  @accept         json
//  @produce        json
//  @param          body    body    DataType    true    "Book form"
//  @success        201
//  @failure        400 {object}    err.Error
//  @failure        422 {object}    err.Errors
//  @failure        500 {object}    err.Error
//  @router         /DataType [post]
func Create(w http.ResponseWriter, r *http.Request) {}

// Read godoc
//
//  @summary        Read all DataType
//  @description    Read all DataType
//  @tags           DataType
//  @accept         json
//  @produce        json
//  @success        200 {array}     DTO
//  @failure        500 {object}    err.Error
//  @router         /DataType [get]
func Read(w http.ResponseWriter, r *http.Request) {}

// Read godoc
//
//  @summary        Read DataType by id
//  @description    Read DataType by id
//  @tags           DataType
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "Book ID"
//  @success        200 {object}    DTO
//  @failure        400 {object}    err.Error
//  @failure        404
//  @failure        500 {object}    err.Error
//  @router         /DataType/{id} [get]
func ReadById(w http.ResponseWriter, r *http.Request) {}

// Update godoc
//
//  @summary        Update DataType
//  @description    Update DataType
//  @tags           DataType
//  @accept         json
//  @produce        json
//  @param          id      path    string  true    "Book ID"
//  @param          body    body    DataType    true    "Book form"
//  @success        200
//  @failure        400 {object}    err.Error
//  @failure        404
//  @failure        422 {object}    err.Errors
//  @failure        500 {object}    err.Error
//  @router         /DataType/{id} [put]
func UpdateById(w http.ResponseWriter, r *http.Request) {}

// Delete godoc
//
//  @summary        Delete DataType
//  @description    Delete DataType
//  @tags           DataType
//  @accept         json
//  @produce        json
//  @param          id  path    string  true    "Book ID"
//  @success        200
//  @failure        400 {object}    err.Error
//  @failure        404
//  @failure        500 {object}    err.Error
//  @router         /DataType/{id} [delete]
func DeleteById(w http.ResponseWriter, r *http.Request) {}

//FIXME: remove this. This just allows importing err for the comments that generate swagger.
func uselessFuncRemoveThis() err.Error {
    e := err.Error{ Error: "lalilulelo" }
    return e
}
