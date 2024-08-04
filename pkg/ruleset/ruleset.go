package ruleset

import (
	"slices"

	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"github.com/ymtdzzz/otel-lint/pkg/rule/attributes"
	"github.com/ymtdzzz/otel-lint/pkg/rule/metrics"
)

type RuleSetVersions map[string]*RuleSet

func (rs RuleSetVersions) RemoveStabilityExperimental() RuleSetVersions {
	for k, rules := range rs {
		// TODO: use interface
		nars := []*rule.AttributeRule{}
		for _, ars := range rules.Attribute {
			if ars.Stability != rule.StabilityExperimental {
				nars = append(nars, ars)
			}
		}
		nmrs := []*rule.MetricRule{}
		for _, mrs := range rules.Metric {
			if mrs.Stability != rule.StabilityExperimental {
				nmrs = append(nmrs, mrs)
			}
		}
		nrs := &RuleSet{
			Attribute: nars,
			Metric:    nmrs,
		}
		rs[k] = nrs
	}

	return rs
}

func (rs RuleSetVersions) RemoveSeverityWarn() RuleSetVersions {
	for k, rules := range rs {
		// TODO: use interface
		nars := []*rule.AttributeRule{}
		for _, ars := range rules.Attribute {
			if ars.Severity != rule.SeverityWarn {
				nars = append(nars, ars)
			}
		}
		nmrs := []*rule.MetricRule{}
		for _, mrs := range rules.Metric {
			if mrs.Severity != rule.SeverityWarn {
				nmrs = append(nmrs, mrs)
			}
		}
		nrs := &RuleSet{
			Attribute: nars,
			Metric:    nmrs,
		}
		rs[k] = nrs
	}

	return rs
}

func (rs RuleSetVersions) RemoveRuleNames(names []string) RuleSetVersions {
	for k, rules := range rs {
		// TODO: use interface
		nars := []*rule.AttributeRule{}
		for _, ars := range rules.Attribute {
			if !slices.Contains(names, ars.Name) {
				nars = append(nars, ars)
			}
		}
		nmrs := []*rule.MetricRule{}
		for _, mrs := range rules.Metric {
			if !slices.Contains(names, mrs.Name) {
				nmrs = append(nmrs, mrs)
			}
		}
		nrs := &RuleSet{
			Attribute: nars,
			Metric:    nmrs,
		}
		rs[k] = nrs
	}

	return rs
}

// Currently supported version is greater than or equal to 1.24.0
// see: https://github.com/open-telemetry/semantic-conventions/tree/main/schemas
func DefaultRuleMap() RuleSetVersions {
	return RuleSetVersions{
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
