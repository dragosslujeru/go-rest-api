package controllers

import (
	"encoding/json"
	"net/http"
)

// GetOnlyHandlerFunc allows only GET requests
type GetOnlyHandlerFunc http.HandlerFunc

func (f GetOnlyHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	f(w, r)
}

// PostOnlyHandlerFunc allows only POST requests
type PostOnlyHandlerFunc http.HandlerFunc

func (f PostOnlyHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	f(w, r)
}

type MethodBasedHandler struct {
	handlersByMethod map[string]http.Handler
}

type HandlerAndMethod struct {
	Method  string
	Handler http.Handler
}

func (h MethodBasedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := h.handlersByMethod[r.Method]
	if !ok {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	handler.ServeHTTP(w, r)
}

//HandleAll will return a handler that will select how to handle request based on method type
func HandleAll(handlers ...HandlerAndMethod) http.Handler {
	handlersByMethod := make(map[string]http.Handler)
	for _, handler := range handlers {
		handlersByMethod[handler.Method] = handler.Handler
	}
	return MethodBasedHandler{handlersByMethod: handlersByMethod}
}

//GetHandler will return a handler and method type pair
func GetHandler(Handler http.Handler) HandlerAndMethod {
	return HandlerAndMethod{http.MethodGet, Handler}
}

//PostHandler will return a handler and method type pair
func PostHandler(Handler http.Handler) HandlerAndMethod {
	return HandlerAndMethod{http.MethodPost, Handler}
}

//WriteJSON writes value as JSON to response
func WriteJSON(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusBadRequest)
	}
	w.Write(json)
}
