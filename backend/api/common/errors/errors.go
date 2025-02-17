package errors
// contains common error messages and useful funcs

import (
	"fmt"
	"net/http"
)

type Error struct {
    error string
}

func InvalidURLParam(param string) (string, int) {
    msg := fmt.Sprintf(`{"error": "invalid url param: %s"}`, param)
    httpStatus := http.StatusBadRequest
    return msg, httpStatus
}

func JsonDecodeFailure() (string, int) {
    msg := fmt.Sprintf(`{"error": "json decode failure"}`)
    httpStatus := http.StatusInternalServerError //FIXME: maybe change this?
    //FIXME: Sometimes (user input) the user is to blame (error 400-something)
    // Sometimes (server read) the server is to blame (error 500-something)
    return msg, httpStatus
}
func JsonEncodeFailure() (string, int) {
    msg := fmt.Sprintf(`{"error": "json encode failure"}`)
    httpStatus := http.StatusInternalServerError
    return msg, httpStatus
}



func DbReadFailure() (string, int) {
    msg := fmt.Sprintf(`{"error": "db data read failure"}`)
    httpStatus := http.StatusInternalServerError
    return msg, httpStatus
}
func DbInsertFailure() (string, int) {
    msg := fmt.Sprintf(`{"error": "db data access failure"}`)
    httpStatus := http.StatusInternalServerError
    return msg, httpStatus
}
func DbUpdateFailure() (string, int) {
    msg := fmt.Sprintf(`{"error": "db data update failure"}`)
    httpStatus := http.StatusInternalServerError
    return msg, httpStatus
}
func DbDeleteFailure() (string, int) {
    msg := fmt.Sprintf(`{"error": "db data delete failure"}`)
    httpStatus := http.StatusInternalServerError
    return msg, httpStatus
}
