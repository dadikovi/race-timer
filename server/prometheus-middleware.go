package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpStatus = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "race_timer_http_status",
		Help: "Statuses of HTTP requests.",
	}, []string{"path", "status"})
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "race_timer_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		d := NewLoggingResponseWriter(w)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(d, r)
		timer.ObserveDuration()
		httpStatus.WithLabelValues(path, fmt.Sprint(d.statusCode)).Inc()
	})
}
