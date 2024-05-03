package metrics

import (
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	metrics.Metrics
	IncNumTasksReceived()
	IncNumTasksAcceptedByAggregator()
	// This metric would either need to be tracked by the aggregator itself,
	// or we would need to write a collector that queries onchain for this info
	// AddPercentageStakeSigned(percentage float64)
}

// AvsMetrics contains instrumented metrics that should be incremented by the avs node using the methods below
type AvsAndEigenMetrics struct {
	metrics.Metrics
	numTasksReceived prometheus.Counter
	// if numSignedTaskResponsesAcceptedByAggregator != numTasksReceived, then there is a bug
	numSignedTaskResponsesAcceptedByAggregator prometheus.Counter
}

const incredibleSquaringNamespace = "incsq"

func NewAvsAndEigenMetrics(avsName string, eigenMetrics *metrics.EigenMetrics, reg prometheus.Registerer) *AvsAndEigenMetrics {
	return &AvsAndEigenMetrics{
		Metrics: eigenMetrics,
		numTasksReceived: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: incredibleSquaringNamespace,
				Name:      "num_tasks_received",
				Help:      "The number of tasks received by reading from the avs service manager contract",
			}),
		numSignedTaskResponsesAcceptedByAggregator: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: incredibleSquaringNamespace,
				Name:      "num_signed_task_responses_accepted_by_aggregator",
				Help:      "The number of signed task responses accepted by the aggregator",
			}),
	}
}

func (m *AvsAndEigenMetrics) IncNumTasksReceived() {
	m.numTasksReceived.Inc()
}

func (m *AvsAndEigenMetrics) IncNumTasksAcceptedByAggregator() {
	m.numSignedTaskResponsesAcceptedByAggregator.Inc()
}
