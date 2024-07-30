package linter

import (
	"net/url"
	"path"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"github.com/ymtdzzz/otel-lint/pkg/ruleset"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var OtelLinter *Linter

func init() {
	OtelLinter = &Linter{
		ruleset: ruleset.RuleMap,
	}
}

type Linter struct {
	ruleset map[string]*ruleset.RuleSet
}

func (l *Linter) RunTrace(pt *ptrace.Traces) (*ptrace.Traces, error) {
	rss := pt.ResourceSpans()
	for rsi := 0; rsi < rss.Len(); rsi++ {
		rs := rss.At(rsi)
		resource := rs.Resource()
		l.runLintAttribute(rs.SchemaUrl(), resource.Attributes())

		for ssi := 0; ssi < rs.ScopeSpans().Len(); ssi++ {
			ss := rs.ScopeSpans().At(ssi)
			scope := ss.Scope()
			l.runLintAttribute(ss.SchemaUrl(), scope.Attributes())

			for si := 0; si < ss.Spans().Len(); si++ {
				span := ss.Spans().At(si)
				l.runLintAttribute(ss.SchemaUrl(), span.Attributes())
			}
		}
	}

	return pt, nil
}

func (l *Linter) RunMetric(pm *pmetric.Metrics) (*pmetric.Metrics, error) {
	rms := pm.ResourceMetrics()
	for rmi := 0; rmi < rms.Len(); rmi++ {
		rm := rms.At(rmi)
		resource := rm.Resource()
		l.runLintAttribute(rm.SchemaUrl(), resource.Attributes())

		for smi := 0; smi < rm.ScopeMetrics().Len(); smi++ {
			sm := rm.ScopeMetrics().At(smi)
			scope := sm.Scope()
			l.runLintAttribute(sm.SchemaUrl(), scope.Attributes())

			for mi := 0; mi < sm.Metrics().Len(); mi++ {
				// metric := sm.Metrics().At(mi)
				// TODO: lint attributes depends on metric type
			}
		}
	}

	return pm, nil
}

func (l *Linter) runLintAttribute(schema string, attrs pcommon.Map) {
	var sb strings.Builder
	semconvVer, err := parseSchemaURL(schema)
	if err != nil {
		sb.WriteString("lint." + rule.SeverityString(rule.SeverityWarn))
		attrs.PutStr(sb.String(), "SchemaURL is not set or invalid format. Checks for this span has been skipped.")
		return
	}
	if _, ok := l.ruleset[semconvVer.String()]; !ok {
		sb.WriteString("lint." + rule.SeverityString(rule.SeverityWarn))
		attrs.PutStr(sb.String(), "Unsupported schema version. Checks for this span has been skipped.")
		return
	}
	sattrs := rule.NewSignalAttributes(attrs)
	var haserr bool
	for _, r := range l.ruleset[semconvVer.String()].Attribute {
		if ok := r.Check(sattrs); !ok {
			haserr = true
			sb.WriteString("lint." + rule.SeverityString(r.Severity) + "." + r.Name)
			attrs.PutStr(sb.String(), r.Title)
		}
		sb.Reset()
	}
	if !haserr {
		attrs.PutStr("lint.result", "OK")
	}
}

func parseSchemaURL(rawURL string) (*semver.Version, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return semver.NewVersion(path.Base(u.Path))
}
