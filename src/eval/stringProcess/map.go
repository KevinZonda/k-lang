package stringProcess

var mapList = map[string]string{
	"\\a":  "\a",
	"\\n":  "\n",
	"\\t":  "\t",
	"\\r":  "\r",
	"\\f":  "\f",
	"\\\\": "\\",
}

var varMap = map[string]string{

	"{{": "{",
	"}}": "}",
}

type Mode int

const (
	ModeNormal Mode = iota
	ModeConst
	ModeVar
)

func hasMapToken(mode Mode, token string) (string, bool) {
	switch mode {
	case ModeNormal:
		val, ok := mapList[token]
		return val, ok
	case ModeConst:
		return token, false
	case ModeVar:
		val, ok := mapList[token]
		if ok {
			return val, true
		}
		val, ok = varMap[token]
		return val, ok
	default:
		return token, false
	}
}
