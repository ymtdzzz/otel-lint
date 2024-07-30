package rule

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type SignalAttributes struct {
	attrs    pcommon.Map
	prefixes map[string]bool
}

func NewSignalAttributes(m pcommon.Map) *SignalAttributes {
	prefixes := map[string]bool{}
	m.Range(func(k string, _ pcommon.Value) bool {
		prefix := strings.Split(k, ".")[0]
		if _, ok := prefixes[prefix]; !ok {
			prefixes[prefix] = true
		}
		return true
	})

	return &SignalAttributes{
		attrs:    m,
		prefixes: prefixes,
	}
}

func (a *SignalAttributes) Get(key string) (pcommon.Value, bool) {
	return a.attrs.Get(key)
}

func (a *SignalAttributes) KeyExists(key string) bool {
	_, ok := a.attrs.Get(key)
	return ok
}

func (a *SignalAttributes) KeyPrefixExists(prefix string) bool {
	_, ok := a.prefixes[prefix]
	return ok
}
