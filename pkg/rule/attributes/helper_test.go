package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymtdzzz/otel-lint/pkg/rule"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestKeyNotExists(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "another.key",
			want:     true,
		},
		{
			name:     "false",
			attrKey:  "test.key",
			inputKey: "test.key",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			pmap.PutStr(tt.attrKey, "value")
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, keyNotExists(attr, tt.inputKey))
		})
	}
}

func TestValueValidEnum(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		enums    []string
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			enums:    []string{"some_value", "value", "another_value"},
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			enums:    []string{"some_value", "value", "another_value"},
			want:     true,
		},
		{
			name:     "false",
			attrKey:  "test.key",
			inputKey: "test.key",
			enums:    []string{"some_value", "wrong_value", "another_value"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			pmap.PutStr(tt.attrKey, "value")
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueValidEnum(attr, tt.inputKey, tt.enums))
		})
	}
}

func TestValueTypeStr(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeStr,
			want:     true,
		},
		{
			name:     "false",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeDouble,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeStr {
				pmap.PutStr(tt.attrKey, "value")
			} else {
				pmap.PutDouble(tt.attrKey, 1234.5)
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeStr(attr, tt.inputKey))
		})
	}
}

func TestValueTypeInt(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeInt,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeInt,
			want:     true,
		},
		{
			name:     "false",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeInt {
				pmap.PutInt(tt.attrKey, 12345)
			} else {
				pmap.PutStr(tt.attrKey, "12345")
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeInt(attr, tt.inputKey))
		})
	}
}

func TestValueTypeDouble(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeDouble,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeDouble,
			want:     true,
		},
		{
			name:     "false",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeDouble {
				pmap.PutDouble(tt.attrKey, 1234.5)
			} else {
				pmap.PutStr(tt.attrKey, "1234.5")
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeDouble(attr, tt.inputKey))
		})
	}
}

func TestValueTypeBool(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeBool,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeBool,
			want:     true,
		},
		{
			name:     "false",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeBool {
				pmap.PutBool(tt.attrKey, true)
			} else {
				pmap.PutStr(tt.attrKey, "true")
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeBool(attr, tt.inputKey))
		})
	}
}

func TestValueTypeStrSlice(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		elmType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeStr,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeStr,
			want:     true,
		},
		{
			name:     "false_value_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
		{
			name:     "false_element_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeDouble,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeSlice {
				sl := pmap.PutEmptySlice(tt.attrKey)
				if tt.elmType == pcommon.ValueTypeStr {
					sl.AppendEmpty().SetStr("value")
				} else {
					sl.AppendEmpty().SetDouble(1234.5)
				}
			} else {
				pmap.PutDouble(tt.attrKey, 1234.5)
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeStrSlice(attr, tt.inputKey))
		})
	}
}

func TestValueTypeIntSlice(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		elmType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeInt,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeInt,
			want:     true,
		},
		{
			name:     "false_value_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
		{
			name:     "false_element_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeSlice {
				sl := pmap.PutEmptySlice(tt.attrKey)
				if tt.elmType == pcommon.ValueTypeInt {
					sl.AppendEmpty().SetInt(12345)
				} else {
					sl.AppendEmpty().SetStr("12345")
				}
			} else {
				pmap.PutStr(tt.attrKey, "12345")
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeIntSlice(attr, tt.inputKey))
		})
	}
}

func TestValueTypeDoubleSlice(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		elmType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeDouble,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeDouble,
			want:     true,
		},
		{
			name:     "false_value_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
		{
			name:     "false_element_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeSlice {
				sl := pmap.PutEmptySlice(tt.attrKey)
				if tt.elmType == pcommon.ValueTypeDouble {
					sl.AppendEmpty().SetDouble(12345)
				} else {
					sl.AppendEmpty().SetStr("1234.5")
				}
			} else {
				pmap.PutStr(tt.attrKey, "1234.5")
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeDoubleSlice(attr, tt.inputKey))
		})
	}
}

func TestValueTypeBoolSlice(t *testing.T) {
	tests := []struct {
		name     string
		attrKey  string
		inputKey string
		valType  pcommon.ValueType
		elmType  pcommon.ValueType
		want     bool
	}{
		{
			name:     "true",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeBool,
			want:     true,
		},
		{
			name:     "true_key_not_exist",
			attrKey:  "test.key",
			inputKey: "another.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeBool,
			want:     true,
		},
		{
			name:     "false_value_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeStr,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
		{
			name:     "false_element_not_match",
			attrKey:  "test.key",
			inputKey: "test.key",
			valType:  pcommon.ValueTypeSlice,
			elmType:  pcommon.ValueTypeStr,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pmap := pcommon.NewMap()
			if tt.valType == pcommon.ValueTypeSlice {
				sl := pmap.PutEmptySlice(tt.attrKey)
				if tt.elmType == pcommon.ValueTypeBool {
					sl.AppendEmpty().SetBool(true)
				} else {
					sl.AppendEmpty().SetStr("true")
				}
			} else {
				pmap.PutStr(tt.attrKey, "true")
			}
			attr := rule.NewSignalAttributes(pmap)
			assert.Equal(t, tt.want, valueTypeBoolSlice(attr, tt.inputKey))
		})
	}
}
