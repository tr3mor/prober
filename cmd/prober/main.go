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

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	var path string
	flag.StringVar(&path, "config", "/etc/prober/config.yaml", "path to config")
	flag.Parse()
	cfg := config.ParseConfig(path, sugar)
	var p []probes.Probe
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", cfg.PrometheusPort), nil)
	}()

	c := http.Client{Timeout: cfg.Timeout}
	for _, target := range cfg.Targets {
		p = append(p, probes.NewHealthcheck(target.Name, sugar, target.URL, &c))
	}
	prober := prober.NewProber(sugar, p, cfg.Period)
	ctx := context.Background()
	prober.Start(ctx)
}
