package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestMetricNameNot(t *testing.T) {
	tests := []struct {
		name      string
		mname     string
		inputName string
		want      bool
	}{
		{
			name:      "true",
			mname:     "test.metric",
			inputName: "another.metric",
			want:      true,
		},
		{
			name:      "false",
			mname:     "test.metric",
			inputName: "test.metric",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := pmetric.NewMetric()
			m.SetName(tt.mname)
			assert.Equal(t, tt.want, metricNameNot(rule.NewMetrics(m), tt.inputName))
		})
	}
}
