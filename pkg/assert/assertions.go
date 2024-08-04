package assert

import (
	"errors"
	"fmt"

	"github.com/ymtdzzz/otel-lint/pkg/linter"
	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
}

type tHelper = interface {
	Helper()
}

const (
	lintResultKeyResource = "Resource"
	lintResultKeySpan     = "Span"
)

func NoSemConvErrorSpan(t TestingT, span trace.ReadOnlySpan) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	res := map[string][]*linter.LintResult{}

	// lint resource
	resource := span.Resource()
	rattrMap, rerr := attributesToMap(resource.Attributes())
	if rerr == nil {
		r := linter.OtelLinter.RunLintAttribute(resource.SchemaURL(), rattrMap)
		if len(r) > 0 {
			res[lintResultKeyResource] = r
		}
	} else {
		res[lintResultKeyResource] = []*linter.LintResult{
			{
				RuleName:  "attribute",
				RuleTitle: fmt.Sprintf("Failed to convert attribute.KeyValue to pcommon.map. error: %v", rerr),
				Severity:  rule.SeverityError,
			},
		}
	}

	// lint span
	sattrMap, serr := attributesToMap(span.Attributes())
	if serr == nil {
		r := linter.OtelLinter.RunLintAttribute(span.InstrumentationScope().SchemaURL, sattrMap)
		if len(r) > 0 {
			res[lintResultKeySpan] = r
		}
	} else {
		res[lintResultKeySpan] = []*linter.LintResult{
			{
				RuleName:  "attribute",
				RuleTitle: fmt.Sprintf("Failed to convert attribute.KeyValue to pcommon.map. error: %v", serr),
				Severity:  rule.SeverityError,
			},
		}
	}

	if len(res) > 0 {
		errmsg := ""
		for k, v := range res {
			m := fmt.Sprintf("%s\n", k)
			for _, e := range v {
				m = fmt.Sprintf("%s- %s\n", m, e.String())
			}
			errmsg = fmt.Sprintf("%s%s\n", errmsg, m)
		}
		return fail(t, errmsg, "Received signal is not conformed to Semantic Conventions")
	}

	return true
}

func attributesToMap(attrs []attribute.KeyValue) (pcommon.Map, error) {
	m := pcommon.NewMap()

	for _, kv := range attrs {
		strkey := string(kv.Key)
		switch kv.Value.Type() {
		case attribute.BOOL:
			m.PutBool(strkey, kv.Value.AsBool())
		case attribute.INT64:
			m.PutInt(strkey, kv.Value.AsInt64())
		case attribute.FLOAT64:
			m.PutDouble(strkey, kv.Value.AsFloat64())
		case attribute.STRING:
			m.PutStr(strkey, kv.Value.AsString())
		case attribute.BOOLSLICE:
			bs := m.PutEmptySlice(strkey)
			for _, b := range kv.Value.AsBoolSlice() {
				bs.AppendEmpty().SetBool(b)
			}
		case attribute.INT64SLICE:
			is := m.PutEmptySlice(strkey)
			for _, i := range kv.Value.AsInt64Slice() {
				is.AppendEmpty().SetInt(i)
			}
		case attribute.FLOAT64SLICE:
			fs := m.PutEmptySlice(strkey)
			for _, f := range kv.Value.AsFloat64Slice() {
				fs.AppendEmpty().SetDouble(f)
			}
		case attribute.STRINGSLICE:
			ss := m.PutEmptySlice(strkey)
			for _, s := range kv.Value.AsStringSlice() {
				ss.AppendEmpty().SetStr(s)
			}
		default:
			return pcommon.NewMap(), errors.New("Invalid attribute type")
		}
	}

	return m, nil
}

// fail reports a failure through
func fail(t TestingT, failureMessage, msg string) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	// Add test name if the Go version supports it
	if n, ok := t.(interface {
		Name() string
	}); ok {
		// TODO: show test name
		n.Name()
	}
	t.Errorf("\n%s\n%s", msg, failureMessage)

	return false
}
