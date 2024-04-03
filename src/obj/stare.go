package obj

func Unbox(from any) (any, bool) {
	if from == nil {
		return nil, false
	}
	switch fromT := from.(type) {
	case *Object:
		return fromT.Value(), true
	}
	return from, false
}

func UnboxToEnd(from any) any {
	switch vT := from.(type) {
	case *Object:
		return UnboxToEnd(vT.Value())
	default:
		return from
	}
}
