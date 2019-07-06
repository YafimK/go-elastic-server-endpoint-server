package controllers

import (
	"fmt"
	"net/http"
)

type BadMethodInRequestError struct {
	expectedMethod string
}

func NewBadMethodInRequestError(expectedMethod string) *BadMethodInRequestError {
	return &BadMethodInRequestError{expectedMethod: expectedMethod}
}

func (e BadMethodInRequestError) Error() string {
	return fmt.Sprintf("bad method in request - expected %v", e.expectedMethod)
}

func Get(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "" {
			http.Error(w, NewBadMethodInRequestError("GET").Error(), http.StatusBadRequest)
		}
		handlerFunc(w, r)
	}
}

func Post(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, NewBadMethodInRequestError("POST").Error(), http.StatusBadRequest)
		}
		handlerFunc(w, r)
	}
}
