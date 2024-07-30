package ruleset

import (
	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"github.com/ymtdzzz/otel-lint/pkg/rule/attributes"
	"github.com/ymtdzzz/otel-lint/pkg/rule/metrics"
)

// Currently supported version is greater than or equal to 1.24.0
// see: https://github.com/open-telemetry/semantic-conventions/tree/main/schemas
var RuleMap = map[string]*RuleSet{
	"1.24.0": {
		Attribute: mergeRules(
			attributes.RulesDeprecatedv1240,
			attributes.RulesTypev1240,
			attributes.RulesEnumv1240,
		),
	},
	"1.25.0": {
		Attribute: mergeRules(
			attributes.RulesDeprecatedv1250,
			attributes.RulesTypev1250,
			attributes.RulesEnumv1250,
		),
	},
	"1.26.0": {
		Attribute: mergeRules(
			attributes.RulesDeprecatedv1260,
			attributes.RulesTypev1260,
			attributes.RulesEnumv1260,
		),
		Metric: metrics.RulesDeprecatedv1260,
	},
}

func mergeRules(rules ...[]*rule.AttributeRule) []*rule.AttributeRule {
	totalLen := 0
	for _, r := range rules {
		totalLen += len(r)
	}

	merged := make([]*rule.AttributeRule, 0, totalLen)

	for _, r := range rules {
		merged = append(merged, r...)
	}

	return merged
}

type RuleSet struct {
	Attribute []*rule.AttributeRule
	// Resource  []*rule.AttributeRule
	// Trace     []*rule.AttributeRule
	Metric []*rule.MetricRule
	// ...
}
