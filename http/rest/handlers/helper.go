package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/erru"
)

func (s service) respond(w http.ResponseWriter, data interface{}, status int) {
	var respData interface{}

	switch v := data.(type) {
	case nil:
	case erru.ErrArgument:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
	default:
		respData = data
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		err := json.NewEncoder(w).Encode(respData)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}
