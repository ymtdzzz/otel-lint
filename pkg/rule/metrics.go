package rule

import "go.opentelemetry.io/collector/pdata/pmetric"

type Metrics struct {
	metrics pmetric.Metric
}

func NewMetrics(m pmetric.Metric) *Metrics {
	return &Metrics{
		metrics: m,
	}
}

func (m *Metrics) Name() string {
	return m.metrics.Name()
}
