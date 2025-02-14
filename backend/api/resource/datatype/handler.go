// FIXME: THIS IS JUST A TEMPLATE
//      this package contains the handlers for the type 'datatype'
//      the handlers are called @api/router/router.go
package datatype

import (
    "net/http"
)

//FIXME: i dont understand why use this kind of type... seems weird af
// apparently I'll attach loggers, validators and Repository to this, giving handler.go (here) access to db etc
//    now, why cant these be used here directly im not exactly sure... manye it is an injection thing (seems to be so)?
//   anyways, need to add func (a *API) to the definition of each handler below later...
// a API instance has additional info for the handlers to work, its like passing it as argument, but this way it can be called as a method of an instace of API created in the scope above... or something like that
type API struct {}

//CRUD handlers for DataType
func Create(w http.ResponseWriter, r *http.Request) {}
func Read(w http.ResponseWriter, r *http.Request) {}
func ReadById(w http.ResponseWriter, r *http.Request) {}
func UpdateById(w http.ResponseWriter, r *http.Request) {}
func DeleteById(w http.ResponseWriter, r *http.Request) {}

