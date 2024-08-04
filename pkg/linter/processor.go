package linter

import (
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func (l *Linter) RunTrace(pt *ptrace.Traces) (*ptrace.Traces, error) {
	rss := pt.ResourceSpans()
	for rsi := 0; rsi < rss.Len(); rsi++ {
		rs := rss.At(rsi)
		resource := rs.Resource()
		l.runLintAttributeWrite(rs.SchemaUrl(), resource.Attributes())
		// TODO: resource-specific lint

		for ssi := 0; ssi < rs.ScopeSpans().Len(); ssi++ {
			ss := rs.ScopeSpans().At(ssi)
			scope := ss.Scope()
			l.runLintAttributeWrite(ss.SchemaUrl(), scope.Attributes())

			for si := 0; si < ss.Spans().Len(); si++ {
				span := ss.Spans().At(si)

				l.runLintAttributeWrite(ss.SchemaUrl(), span.Attributes())
				// TODO: span-specific lint
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
		l.runLintAttributeWrite(rm.SchemaUrl(), resource.Attributes())
		// TODO: resource-specific lint

		for smi := 0; smi < rm.ScopeMetrics().Len(); smi++ {
			sm := rm.ScopeMetrics().At(smi)
			scope := sm.Scope()
			l.runLintAttributeWrite(sm.SchemaUrl(), scope.Attributes())

			for mi := 0; mi < sm.Metrics().Len(); mi++ {
				// metric := sm.Metrics().At(mi)
				// TODO: lint attributes depends on metric type
				// TODO: metric-specific lint
			}
		}
	}

	return pm, nil
}

func (l *Linter) RunLog(pl *plog.Logs) (*plog.Logs, error) {
	rls := pl.ResourceLogs()
	for rli := 0; rli < rls.Len(); rli++ {
		rl := rls.At(rli)
		resource := rl.Resource()
		l.runLintAttributeWrite(rl.SchemaUrl(), resource.Attributes())
		// TODO: resource-specific lint

		for sli := 0; sli < rl.ScopeLogs().Len(); sli++ {
			sl := rl.ScopeLogs().At(sli)
			scope := sl.Scope()
			l.runLintAttributeWrite(sl.SchemaUrl(), scope.Attributes())

			for li := 0; li < sl.LogRecords().Len(); li++ {
				// log := sl.LogRecords().At(li)
				// TODO: log-specific lint
			}
		}
	}

	return pl, nil
}

func (l *Linter) runLintAttributeWrite(schema string, attrs pcommon.Map) {
	res := l.RunLintAttribute(schema, attrs)
	if len(res) == 0 {
		attrs.PutStr("lint.result", "OK")
		return
	}

	for _, v := range res {
		attrs.PutStr(fmt.Sprintf("lint.%s.%s", v.Severity.String(), v.RuleName), v.RuleTitle)
	}
}
