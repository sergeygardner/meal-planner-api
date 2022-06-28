package service

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	"net/http"
	"runtime"
	"runtime/debug"
	"strconv"
)

func Error400HandleService(w http.ResponseWriter, err error) *response.Error {
	return ErrorHandleService(http.StatusBadRequest, w, err)
}

func ErrorHandleService(status int, w http.ResponseWriter, err error) *response.Error {
	w.WriteHeader(status)
	id, _ := uuid.NewUUID()
	_, _, line, _ := runtime.Caller(1)
	return &response.Error{Id: id, Status: status, Line: line, Code: strconv.Itoa(status), Message: err.Error(), Trace: string(debug.Stack())}
}
