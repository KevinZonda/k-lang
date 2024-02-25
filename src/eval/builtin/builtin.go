package builtin

type FuncCallResult struct {
	hasValue bool
	value    any
}

func (f FuncCallResult) HasValue() bool {
	return f.hasValue
}

func (f FuncCallResult) Value() any {
	return f.value
}

func resultVal(v any) FuncCallResult {
	return FuncCallResult{hasValue: true, value: v}
}

func resultNoVal() FuncCallResult {
	return FuncCallResult{hasValue: false, value: nil}
}
