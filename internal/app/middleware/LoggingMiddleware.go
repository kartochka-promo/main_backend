package middleware

import (
	"2020_1_drop_table/internal/pkg/metrics"
	"2020_1_drop_table/internal/pkg/responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"time"
)

type LoggingResponseWriter struct {
	wrapped  http.ResponseWriter
	response responses.HttpResponse
}

func NewLoggingResponseWriter(wrapped http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{wrapped: wrapped}
}

func (lrw *LoggingResponseWriter) Header() http.Header {
	return lrw.wrapped.Header()
}

func (lrw *LoggingResponseWriter) Write(content []byte) (int, error) {
	// error ignored because there are cases, where response is OK and body is nil(or contains photo)
	_ = json.Unmarshal(content, &lrw.response)
	return lrw.wrapped.Write(content)
}

func (lrw *LoggingResponseWriter) WriteHeader(statusCode int) {
	lrw.wrapped.WriteHeader(statusCode)
}

func NewLoggingMiddleware(metrics *metrics.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/metrics" {
				msg := fmt.Sprintf("URL: %s, METHOD: %s", r.RequestURI, r.Method)
				log.Info().Msgf(msg)
			}

			reqTime := time.Now()
			w2 := NewLoggingResponseWriter(w)
			next.ServeHTTP(w2, r)
			respTime := time.Since(reqTime)
			if r.URL.Path != "/metrics" {
				if w2.response.Errors == nil || len(w2.response.Errors) == 0 {
					metrics.Hits.WithLabelValues(strconv.Itoa(http.StatusOK), r.URL.Path, r.Method).Inc()
				} else {
					metrics.Hits.WithLabelValues(strconv.Itoa(w2.response.Errors[0].Code), r.URL.Path, r.Method).Inc()
				}
				metrics.Timings.WithLabelValues(
					strconv.Itoa(http.StatusOK), r.URL.String(), r.Method).Observe(respTime.Seconds())
			}
		})
	}
}
