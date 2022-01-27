package probes

import (
	"go.uber.org/zap"
	"net/http"
	"prober/internal/app/prober/metric"
	"time"
)

type Healthcheck struct {
	name   string
	l      *zap.SugaredLogger
	url    string
	client *http.Client
	M      *metric.ProbeMetrics
}

func NewHealthcheck(name string, l *zap.SugaredLogger, url string, c *http.Client) *Healthcheck {
	return &Healthcheck{name, l, url, c, metric.NewMetric(name)}
}

func (h *Healthcheck) Name() string {
	return h.name

}

func (h *Healthcheck) Execute() {
	h.l.With("probe", h.name)
	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		h.l.Errorf("Failed to create http request: %s", err)
		return
	}
	start := time.Now()
	resp, err := h.client.Do(req)
	if err != nil {
		h.l.Errorf("Failed to execute http request: %s", err)
		return
	}
	var status float64
	if resp.StatusCode == http.StatusOK {
		status = 1
	} else {
		status = 0
	}
	duration := time.Since(start).Seconds()
	h.M.CollectMetrics(status, duration)
}
