package simple

import (
	"context"
	"testing"

	// "github.com/ymtdzzz/otel-lint/pkg/assert"
	// "github.com/ymtdzzz/otel-lint/pkg/linter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestDoTrace(t *testing.T) {
	sr := tracetest.NewSpanRecorder()
	tp := trace.NewTracerProvider(trace.WithSpanProcessor(sr))
	otel.SetTracerProvider(tp)

	doTrace(context.Background())

	/*
		for _, s := range sr.Ended() {
			assert.NoSemConvErrorSpan(t, s, linter.IgnoreRules([]string{"deprecated.http.user_agent"}))
		}
	*/
}
