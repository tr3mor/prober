package probes

type Probe interface {
	Execute()
	Name() string
}
