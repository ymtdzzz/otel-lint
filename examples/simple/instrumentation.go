package simple

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

func doTrace(ctx context.Context) {
	tracer := otel.GetTracerProvider().Tracer("example.io/instrumentation/simple", trace.WithSchemaURL(semconv.SchemaURL))
	attrs := []attribute.KeyValue{
		attribute.String("http.method", "GET"),
		attribute.String("http.user_agent", "hoge"),
		attribute.Int("http.host", 1234),
	}
	_, c := tracer.Start(ctx, "hoge", trace.WithAttributes(attrs...), trace.WithSpanKind(trace.SpanKindClient))
	defer c.End()
}
