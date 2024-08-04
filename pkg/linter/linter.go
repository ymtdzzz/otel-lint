package linter

import (
	"fmt"
	"net/url"
	"path"

	"github.com/Masterminds/semver/v3"
	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"github.com/ymtdzzz/otel-lint/pkg/ruleset"
)

var DefaultLinter = Linter{
	ruleset: ruleset.DefaultRuleMap(),
}

type LintResult struct {
	RuleName    string
	RuleTitle   string
	RuleVersion string
	Severity    rule.Severity
}

func (r *LintResult) String() string {
	return fmt.Sprintf("[%s] %s (%s) (semconv version: %s)", r.Severity.String(), r.RuleTitle, r.RuleName, r.RuleVersion)
}

type Linter struct {
	ruleset ruleset.RuleSetVersions
}

func NewLinter(options ...Option) *Linter {
	l := DefaultLinter

	for _, opt := range options {
		opt(&l)
	}

	return &l
}

func parseSchemaURL(rawURL string) (*semver.Version, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return semver.NewVersion(path.Base(u.Path))
}
