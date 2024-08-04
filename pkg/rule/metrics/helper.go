package metrics

import "github.com/ymtdzzz/otel-lint/pkg/rule"

func metricNameNot(m *rule.Metrics, name string) bool {
	return m.Name() != name
}
