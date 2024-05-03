package metrics

import (
	eigenmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
)

type NoopMetrics struct {
	eigenmetrics.NoopMetrics
}

func NewNoopMetrics() *NoopMetrics {
	return &NoopMetrics{
		NoopMetrics: *eigenmetrics.NewNoopMetrics(),
	}
}

func (m *NoopMetrics) IncNumTasksReceived() {}

func (m *NoopMetrics) IncNumTasksAcceptedByAggregator() {}
