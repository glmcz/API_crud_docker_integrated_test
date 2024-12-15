package http

import (
	"net/http"
)

type HttpResponse struct {
	Code int
	Msg  string
}

func (h *HttpResponse) WriteResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Code)
	_, err := w.Write([]byte(h.Msg))
	if err != nil {
		return
	}
}
