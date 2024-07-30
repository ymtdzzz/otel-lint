package rule

type Severity int

const (
	SeverityError = iota
	SeverityWarn
	SeverityInfo
)

func SeverityString(s Severity) string {
	switch s {
	case SeverityError:
		return "error"
	case SeverityWarn:
		return "warn"
	case SeverityInfo:
		return "info"
	default:
		return "unknown"
	}
}
