package handlers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
}

type messageHandler interface {
	getMessage() interface{}
}

type Handlers struct {
	*http.ServeMux
	Links []string
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	return &Handlers{
		ServeMux: mux,
		Links:    make([]string, 0),
	}
}

func (hh *Handlers) AddHandler(path string, h messageHandler) *Handlers {
	hh.Links = append(hh.Links, path)
	hh.ServeMux.HandleFunc(path, getHandler(h))
	return hh
}

func getHandler(ms messageHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := response{
			Message: ms.getMessage(),
			Status:  200,
		}
		enc := json.NewEncoder(w)
		enc.SetIndent("", "   ")
		enc.Encode(res)
	}
}
