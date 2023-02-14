package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"prober/internal/app/prober"
	"prober/internal/app/prober/config"

	"net/http"
	"prober/internal/app/prober/probes"
)

func k8sProbe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData := []byte(`{"status":"OK"}`)
	_, err := w.Write(jsonData)
	if err != nil {
		return
	}
}

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	sugar := logger.Sugar()
	var path string
	flag.StringVar(&path, "config", "/etc/prober/config.yaml", "path to config")
	flag.Parse()
	cfg := config.ParseConfig(path, sugar)
	var p []probes.Probe
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", k8sProbe)
	http.HandleFunc("/ready", k8sProbe)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PrometheusPort), nil)
		if err != nil {
			sugar.Errorf("Failed to start http server due to: %s", err)
		}
	}()

	c := http.Client{Timeout: cfg.Timeout}
	for _, target := range cfg.Targets {
		p = append(p, probes.NewHealthcheck(target.Name, sugar, target.URL, &c))
	}
	prob := prober.NewProber(sugar, p, cfg.Period)
	ctx := context.Background()
	prob.Start(ctx)
}
