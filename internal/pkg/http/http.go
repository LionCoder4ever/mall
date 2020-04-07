package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
	"mall/internal/pkg/ecode"
	"net/http"
)

const Apiversion = "/v1"

var jsonContentType = []string{"application/json; charset=utf-8"}

// http
type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func writeJSON(w http.ResponseWriter, data interface{}) (err error) {
	var jsonBytes []byte
	writeContentType(w, jsonContentType)
	if jsonBytes, err = json.Marshal(data); err != nil {
		err = errors.WithStack(err)
		return
	}
	if _, err = w.Write(jsonBytes); err != nil {
		err = errors.WithStack(err)
	}
	return
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func (r Response) Render(w http.ResponseWriter) (err error) {
	return writeJSON(w, r)
}

// WriteContentType write json ContentType.
func (r Response) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

/**
transfer err code to msg
*/
func WrapResponse(data interface{}, err error) (int, render.Render) {
	code := http.StatusOK
	customCode := ecode.Cause(err)
	return code, Response{
		Code:    customCode.Code(),
		Message: customCode.Message(),
		Data:    data,
	}
}
