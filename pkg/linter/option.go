package linter

type Option func(*Linter)

func IgnoreExperimental() Option {
	return func(l *Linter) {
		l.ruleset = l.ruleset.RemoveStabilityExperimental()
	}
}

func IgnoreWarn() Option {
	return func(l *Linter) {
		l.ruleset = l.ruleset.RemoveSeverityWarn()
	}
}

func IgnoreRules(rules []string) Option {
	return func(l *Linter) {
		l.ruleset = l.ruleset.RemoveRuleNames(rules)
	}
}
