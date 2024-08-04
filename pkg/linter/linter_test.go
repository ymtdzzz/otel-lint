package linter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymtdzzz/otel-lint/pkg/rule/attributes"
	"github.com/ymtdzzz/otel-lint/pkg/ruleset"
	"github.com/ymtdzzz/otel-lint/test"
)

func TestRunTrace(t *testing.T) {
	t.Run("Schema_Check", func(t *testing.T) {
		pt, ss := test.GenerateOTLPTracesPayload(t, 3, 1, 1)
		ss.RSpans[0].SetSchemaUrl("invalid")
		ss.RSpans[1].SetSchemaUrl("https://opentelemetry.io/schemas/1.1.0")
		ss.RSpans[2].SetSchemaUrl("https://opentelemetry.io/schemas/1.26.0")
		l := &Linter{
			ruleset: map[string]*ruleset.RuleSet{
				"1.26.0": {
					Attribute: attributes.RulesDeprecatedv1260,
				},
			},
		}
		res, err := l.RunTrace(pt)
		assert.Nil(t, err)

		// Resource 1: Invalid schema
		val, ok := res.ResourceSpans().At(0).Resource().Attributes().Get("lint.warn.schema_url")
		assert.True(t, ok)
		assert.Equal(t, "SchemaURL is not set or invalid format. Checks for this span has been skipped.", val.AsString())

		// Resource 2: Invalid schema version
		val, ok = res.ResourceSpans().At(1).Resource().Attributes().Get("lint.warn.schema_version")
		assert.True(t, ok)
		assert.Equal(t, "Unsupported schema version. Checks for this span has been skipped.", val.AsString())

		// Resource 3: OK
		_, ok = res.ResourceSpans().At(2).Resource().Attributes().Get("lint.warn")
		assert.False(t, ok)
	})

	t.Run("Rule_Check", func(t *testing.T) {
		pt, ss := test.GenerateOTLPTracesPayload(t, 2, 1, 1)
		ss.RSpans[0].SetSchemaUrl("https://opentelemetry.io/schemas/1.26.0")
		ss.RSpans[0].Resource().Attributes().PutStr("container.labels", "my-label")
		ss.RSpans[0].Resource().Attributes().PutStr("db.instance.id", "my-id")
		ss.RSpans[1].SetSchemaUrl("https://opentelemetry.io/schemas/1.26.0")
		ss.RSpans[1].Resource().Attributes().PutStr("container.labels", "my-label")
		l := &Linter{
			ruleset: map[string]*ruleset.RuleSet{
				"1.26.0": {
					Attribute: attributes.RulesDeprecatedv1260,
				},
			},
		}
		res, err := l.RunTrace(pt)
		assert.Nil(t, err)

		// Resource 1: Multiple rules
		val, ok := res.ResourceSpans().At(0).Resource().Attributes().Get("lint.result")
		assert.False(t, ok)
		val, ok = res.ResourceSpans().At(0).Resource().Attributes().Get("lint.error.deprecated.container.labels")
		assert.True(t, ok)
		assert.Equal(t, "Deprecated, use `container.label` instead.", val.AsString())
		val, ok = res.ResourceSpans().At(0).Resource().Attributes().Get("lint.error.deprecated.db.instance.id")
		assert.True(t, ok)
		assert.Equal(t, "Deprecated, no general replacement at this time. For Elasticsearch, use `db.elasticsearch.node.name` instead.", val.AsString())

		// Resource 2: Single rule
		val, ok = res.ResourceSpans().At(1).Resource().Attributes().Get("lint.result")
		assert.False(t, ok)
		val, ok = res.ResourceSpans().At(1).Resource().Attributes().Get("lint.error.deprecated.container.labels")
		assert.True(t, ok)
		assert.Equal(t, "Deprecated, use `container.label` instead.", val.AsString())
		_, ok = res.ResourceSpans().At(1).Resource().Attributes().Get("lint.error.R1002")
		assert.False(t, ok)
	})
}
