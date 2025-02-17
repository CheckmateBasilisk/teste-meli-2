package handlers

import (
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/go-chi/chi/v5"
)


// fill in data type
type DataType struct {
    Foo int `json:"foo"`
    Bar string `json:"bar"`
}

var mockDataType = []DataType{
    {0,"000"},
    {1,"001"},
    {2,"002"},
    {3,"003"},
    {4,"004"},
    {5,"005"},
    {6,"006"},
}

//CRUD handlers for DataType
func CreateDataType(w http.ResponseWriter, r *http.Request) {
    // retrieve data from body
    var data DataType
    json.NewDecoder(r.Body).Decode(&data) //FIXME: add error treatment
    //add to db
    mockDataType = append(mockDataType, data)
}
func ReadDataType(w http.ResponseWriter, r *http.Request) {
    // retrieve data
    result := mockDataType
    // build response
    response, _ := json.Marshal(result) //FIXME: change this to newEncoder ?//FIXME: add error treatment
    w.Write([]byte(string(response)))
}
func ReadDataTypeById(w http.ResponseWriter, r *http.Request) {
    // retrieve data
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))//FIXME: add error treatment
    // error treatment
    result := mockDataType[id]
    // build response
    response, _ := json.Marshal(result)//FIXME: add error treatment
    w.Write([]byte(string(response)))
}
func UpdateDataTypeById(w http.ResponseWriter, r *http.Request) {
    // retrieve data
    id, _ := strconv.Atoi(chi.URLParam(r, "id")) //FIXME: add error treatment
    var newData DataType
    json.NewDecoder(r.Body).Decode(&newData) //FIXME: add error treatment
    //error handling
    //add to db
    mockDataType[id] = newData
}
func DeleteDataTypeById(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))//FIXME: add error treatment

    //remove from db
    mockDataType = append(mockDataType[:id], mockDataType[id+1:]...)
}
