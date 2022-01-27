package prober

import (
	"context"
	"go.uber.org/zap"
	probes2 "prober/internal/app/prober/probes"
	"sync"
	"time"
)

type Prober struct {
	log    *zap.SugaredLogger
	probes []probes2.Probe
	period time.Duration
}

func NewProber(l *zap.SugaredLogger, p []probes2.Probe, period time.Duration) *Prober {
	return &Prober{l, p, period}
}

func (p *Prober) Start(ctx context.Context) {
	p.executeProbes(ctx)
	ticker := time.NewTicker(p.period)
	for {
		select {
		case <-ticker.C:
			p.executeProbes(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (p *Prober) executeProbes(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(len(p.probes))
	for _, probe := range p.probes {
		p.log.With("probe", probe.Name())
		go func(ctx context.Context, probe probes2.Probe) {
			defer wg.Done()
			probe.Execute()
		}(ctx, probe)
	}
}
