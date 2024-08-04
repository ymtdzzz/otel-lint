package attributes

import (
	"slices"

	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func keyNotExists(a *rule.SignalAttributes, key string) bool {
	return !a.KeyExists(key)
}

func valueValidEnum(a *rule.SignalAttributes, key string, enums []string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	return slices.Contains(enums, val.Str())
}

func valueTypeStr(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	return val.Type() == pcommon.ValueTypeStr
}

func valueTypeInt(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	return val.Type() == pcommon.ValueTypeInt
}

func valueTypeDouble(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	return val.Type() == pcommon.ValueTypeDouble
}

func valueTypeBool(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	return val.Type() == pcommon.ValueTypeBool
}

func valueTypeStrSlice(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	if val.Type() != pcommon.ValueTypeSlice {
		return false
	}
	for _, v := range val.Slice().AsRaw() {
		if _, ok := v.(string); !ok {
			return false
		}
	}
	return true
}

func valueTypeIntSlice(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	if val.Type() != pcommon.ValueTypeSlice {
		return false
	}
	for _, v := range val.Slice().AsRaw() {
		if _, ok := v.(int64); !ok {
			return false
		}
	}
	return true
}

func valueTypeDoubleSlice(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	if val.Type() != pcommon.ValueTypeSlice {
		return false
	}
	for _, v := range val.Slice().AsRaw() {
		if _, ok := v.(float64); !ok {
			return false
		}
	}
	return true
}

func valueTypeBoolSlice(a *rule.SignalAttributes, key string) bool {
	val, ok := a.Get(key)
	if !ok {
		return true
	}
	if val.Type() != pcommon.ValueTypeSlice {
		return false
	}
	for _, v := range val.Slice().AsRaw() {
		if _, ok := v.(bool); !ok {
			return false
		}
	}
	return true
}
