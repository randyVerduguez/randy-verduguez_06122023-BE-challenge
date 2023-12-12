package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func (s service) MiddlewareLogger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/healthz" {
				next.ServeHTTP(w, r)
			}

			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			requestBodyBytes, err := s.readRequestBody(r)

			if err != nil {
				s.respond(w, err, http.StatusInternalServerError)
				return
			}

			s.restoreRequestBody(r, requestBodyBytes)

			timestamp := time.Now().UTC().Format("2006-01-02 15:04:05 MST")

			regex := regexp.MustCompile("[\n\r\"]")
			payload := string(requestBodyBytes)
			payload = regex.ReplaceAllString(payload, "")

			endpointURL := fmt.Sprintf("%s %s%s", r.Method, r.Host, r.RequestURI)

			// timestamp | payload | endpoint url
			logMessage := fmt.Sprintf("%s | %s | %s", timestamp, payload, endpointURL)

			next.ServeHTTP(w, r)

			s.logger.Trace(logMessage)
		}

		return http.HandlerFunc(fn)
	}
}
