package http_server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"l0_wb/internal/utils/logger"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wdata := []any{
			"method", r.Method,
			"url", r.URL,
		}
		if r.Method == http.MethodPost {
			wdata = append(wdata, "body")

			body, err := io.ReadAll(r.Body)
			var bodyMsg string
			if err != nil {
				bodyMsg = fmt.Sprintf("error: body unavailable: %v", err)
			} else {
				bodyMsg = string(body)
			}

			wdata = append(wdata, bodyMsg)
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		logger.Log().Infow("HTTP_REQ", wdata...)

		next.ServeHTTP(w, r)
	})
}
