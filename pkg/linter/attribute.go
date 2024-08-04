package linter

import (
	"sort"

	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (l *Linter) RunLintAttribute(schema string, attrs pcommon.Map) []*LintResult {
	res := []*LintResult{}
	semconvVer, err := parseSchemaURL(schema)
	if err != nil {
		return append(res, &LintResult{
			RuleName:    "schema_url",
			RuleTitle:   "SchemaURL is not set or invalid format. Checks for this span has been skipped.",
			RuleVersion: "N/A",
			Severity:    rule.SeverityWarn,
		})
	}
	if _, ok := l.ruleset[semconvVer.String()]; !ok {
		return append(res, &LintResult{
			RuleName:    "schema_version",
			RuleTitle:   "Unsupported schema version. Checks for this span has been skipped.",
			RuleVersion: "N/A",
			Severity:    rule.SeverityWarn,
		})
	}
	sattrs := rule.NewSignalAttributes(attrs)
	for _, r := range l.ruleset[semconvVer.String()].Attribute {
		if ok := r.Check(sattrs); !ok {
			res = append(res, &LintResult{
				RuleName:    r.Name,
				RuleTitle:   r.Title,
				RuleVersion: r.Version,
				Severity:    r.Severity,
			})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].RuleName < res[j].RuleName
	})

	return res
}
