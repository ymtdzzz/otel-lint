package linter

type Option func(*Linter)

func WithoutExperimental() Option {
	return func(l *Linter) {
		l.ruleset = l.ruleset.RemoveStabilityExperimental()
	}
}

func WithoutWarn() Option {
	return func(l *Linter) {
		l.ruleset = l.ruleset.RemoveSeverityWarn()
	}
}

func WithoutRules(rules []string) Option {
	return func(l *Linter) {
		l.ruleset = l.ruleset.RemoveRuleNames(rules)
	}
}
