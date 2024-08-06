# otel-lint

otel-lint is a library to check whether OpenTelemetry instrumentation libraries adhere to Semantic Conventions. It provides a linting feature to ensure your instrumentation is up to standard.

## Features

### Integrations

- **`assert` package for Golang**: Lint checks for your instrumentation library development flow
- **Custom processor `lintprocessor`**: End-to-end lint checks for your instrumented application's pipeline
  - see: https://github.com/ymtdzzz/opentelemetry-collector-extra/tree/main/processor/lintprocessor

### Supported Rules

Currently, this tool supports Semantic Convention versions from `v1.24.0` to `v1.27.0`.

| Type      | Rule Name      | Rule ID                          | Severity | Description                                 |
| --------- | -------------- | -------------------------------- | -------- | ------------------------------------------- |
| \-        | Schema URL     | `schema_url\`                    | Warn     | SchemaURL is not set or invalid format      |
| \-        | Schema Version | `schema_version\`                | Warn     | Unsupported schema version                  |
| Attribute | Deprecated     | `deprecated.{{ attribute key }}` | Error    | Deprecated attribute key is used            |
| Attribute | Type           | `type.{{ attribute_key }}`       | Error    | Attribute value is wrong                    |
| Attribute | Enum           | `enum.{{ attribute_key }}`       | Error    | Attribute value is not allowed in enum list |
| Resource  | \-             | \-                               | \-       | Not yet supported                           |
| Trace     | \-             | \-                               | \-       | Not yet supported                           |
| Metric    | Deprecated     | `deprecated.{{ metric name }}`   | Error    | Deprecated metric name is used              |
| Log       | \-             | -                                | \-       | Not yet supported                           |

## Getting Started
### For Golang Instrumentation Libraries

Add this package to `go.mod`.

```sh
$ go get -u github.com/ymtdzzz/otel-lint@latest
```

Use the `assert` and `linter` packages to write assertions for your instrumentation.

```go
package simple

import (
	"context"
	"testing"

	"github.com/ymtdzzz/otel-lint/pkg/assert"
	"github.com/ymtdzzz/otel-lint/pkg/linter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestDoTrace(t *testing.T) {
	sr := tracetest.NewSpanRecorder()
	tp := trace.NewTracerProvider(trace.WithSpanProcessor(sr))
	otel.SetTracerProvider(tp)

	doTrace(context.Background())

	for _, s := range sr.Ended() {
		assert.NoSemConvErrorSpan(t, s, linter.IgnoreRules([]string{"deprecated.http.user_agent"}))
	}
}
```

You can get pretty printed fail messages if any lint issues are detected.

```sh
?   	github.com/ymtdzzz/otel-lint/pkg/assert	[no test files]
=== RUN   TestDoTrace
    instrumentation_test.go:22:
        Received signal is not conformed to Semantic Conventions
        Span
        - [error] Deprecated, use one of `server.address`, `client.address` or `http.request.header.host` instead, depending on the usage. (deprecated.http.host) (semconv version: v1.26.0)
        - [error] Deprecated, use `http.request.method` instead. (deprecated.http.method) (semconv version: v1.26.0)
        - [error] The type of `http.host` should be string (type.http.host) (semconv version: v1.26.0)

--- FAIL: TestDoTrace (0.00s)
FAIL
FAIL	github.com/ymtdzzz/otel-lint/examples/simple	0.005s
```

For more examples, see the `/examples` directory. You can check all lint options in `/pkg/linter/option.go`

### For OpenTelemetry Collector Pipelines (with lintprocessor)

see: https://github.com/ymtdzzz/opentelemetry-collector-extra/tree/main/processor/lintprocessor

TBD

## Configuration

TBD

## Contributing
### Development

TBD

## License

TBD
