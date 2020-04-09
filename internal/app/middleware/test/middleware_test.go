package test

import (
	middleware2 "2020_1_drop_table/internal/app/middleware"
	"context"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPanic(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test")
		assert.True(t, true)
	})
	handlerToTest := middleware2.PanicMiddleware(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)
	session := sessions.Session{Values: map[interface{}]interface{}{"userID": 228}}
	req = req.WithContext(context.WithValue(req.Context(), "session", &session))

	recorder := httptest.NewRecorder()

	handlerToTest.ServeHTTP(recorder, req)
}

func TestLog(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	})
	handlerToTest := middleware2.LoggingMiddleware(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)
	session := sessions.Session{Values: map[interface{}]interface{}{"userID": 228}}
	req = req.WithContext(context.WithValue(req.Context(), "session", &session))

	recorder := httptest.NewRecorder()
	handlerToTest.ServeHTTP(recorder, req)
}
