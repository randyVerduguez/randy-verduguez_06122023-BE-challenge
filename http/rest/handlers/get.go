package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/http/rest/types"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/erru"
)

func (s service) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request types.Request

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("Invalid request body"),
			}, http.StatusBadRequest)
			return
		}

		if request == (types.Request{}) {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("Invalid request body"),
			}, http.StatusBadRequest)
			return
		}

		params := reflect.ValueOf(request)

		for i := 0; i < params.NumField(); i++ {
			param := params.Field(i).String()
			if len(param) == 0 || !isAlpha(param) {
				s.respond(w, erru.ErrArgument{
					Wrapped: errors.New("Invalid request body"),
				}, http.StatusBadRequest)
				return
			}
		}

		request.Name = strings.TrimSpace(strings.ToLower(request.Name))
		request.Region = strings.TrimSpace(strings.ToLower(request.Region))
		request.Country = strings.TrimSpace(strings.ToLower(request.Country))
		getResponse, err := s.toDoService.Get(r.Context(), request)

		if err != nil {
			s.respond(w, err, http.StatusBadRequest)
			return
		}

		s.respond(w, types.Response{
			Name:       getResponse.Name,
			Region:     getResponse.Region,
			Country:    getResponse.Country,
			TempC:      getResponse.TempC,
			TempF:      getResponse.TempF,
			FeelsLikeC: getResponse.FeelsLikeC,
			FeelsLikeF: getResponse.FeelsLikeF,
			WindMph:    getResponse.WindMph,
			WindKph:    getResponse.WindKph,
			Humidity:   getResponse.Humidity,
			CreatedOn:  getResponse.CreatedOn,
		}, http.StatusOK)
	}
}

func isAlpha(s string) bool {
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}

		if !unicode.IsLetter(c) {
			return false
		}
	}

	return true
}
