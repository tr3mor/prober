package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type ProbeMetrics struct {
	Name     string
	Status   prometheus.Gauge
	Duration prometheus.Gauge
}

func NewMetric(name string) *ProbeMetrics {
	key := strings.ReplaceAll(name, "-", "_")
	m := &ProbeMetrics{
		name,
		prometheus.NewGauge(prometheus.GaugeOpts{Name: fmt.Sprintf("%s_probe_status", key)}),
		prometheus.NewGauge(prometheus.GaugeOpts{Name: fmt.Sprintf("%s_probe_duration_seconds", key)}),
	}
	prometheus.MustRegister(m.Status, m.Duration)
	return m
}
func (m *ProbeMetrics) CollectMetrics(status float64, duration float64) {
	m.Status.Set(status)
	m.Duration.Set(duration)
}
