package rule

type AttributeRule struct {
	Name      string
	Title     string
	Check     func(*SignalAttributes) bool
	Severity  Severity
	Stability Stability
	Source    string
}

type MetricRule struct {
	Name      string
	Title     string
	Check     func(*Metrics) bool
	Severity  Severity
	Stability Stability
	Source    string
}
