package core

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var dbDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "race_timer_sql_duration_seconds",
	Help: "Duration of SQL queries.",
}, []string{"query"})

func startDbTimer(op string) *prometheus.Timer {
	return prometheus.NewTimer(dbDuration.WithLabelValues(op))
}
