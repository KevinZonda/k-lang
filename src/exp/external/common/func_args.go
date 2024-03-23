package common

type FuncArgs []ArgType

type ArgType string

const (
	ArgTypeInt ArgType = "int"
	ArgTypeF64 ArgType = "float64"
	ArgTypeStr ArgType = "string"
	ArgTypeArr ArgType = "[]"
	ArgTypeMap ArgType = "map"
	ArgTypeBoo ArgType = "bool"
)

func ConvertArg(v []any, t []ArgType) []any {
	if len(v) != len(t) {
		panic("Length mismatch")
	}
	var r []any
	for i, vT := range v {
		r = append(r, t[i].From(vT))
	}
	return r
}

func (argType ArgType) From(v any) any {
	switch argType {
	case ArgTypeInt:
		switch vT := v.(type) {
		case int:
			return vT
		case int64:
			return int(vT)
		case float64:
			return int(vT)
		case float32:
			return int(vT)
		}
	case ArgTypeF64:
		switch vT := v.(type) {
		case float64:
			return vT
		case int:
			return float64(vT)
		case int64:
			return float64(vT)
		}
	case ArgTypeStr:
		switch vT := v.(type) {
		case string:
			return vT
		}
	case ArgTypeBoo:
		switch vT := v.(type) {
		case bool:
			return vT
		}
	case ArgTypeArr:
		switch vT := v.(type) {
		case []any:
			return vT
		}
	case ArgTypeMap:
		switch vT := v.(type) {
		case map[any]any:
			return vT
		}

	}
	panic("Cannot convert")
}
