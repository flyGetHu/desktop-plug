package common

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Extra interface{} `json:"extra"`
}

func (r Result) Success(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	r.Code = http.StatusOK
	e.Encode(r)
}

func (r Result) Error(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	r.Code = http.StatusInternalServerError
	if r.Msg == "" {
		r.Msg = "操作失败,稍后再试!"
	}
	e := json.NewEncoder(w)
	r.Code = http.StatusOK
	e.Encode(r)
}
