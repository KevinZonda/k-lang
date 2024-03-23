package common

import (
	"encoding/json"
)

type PackageInfoElement struct {
	Type ElementType `json:"type"`
	Args []ArgType   `json:"args"`
}

type ElementType string

const (
	Function ElementType = "func"
	Variable ElementType = "var"
)

type InvokeResult struct {
	Success   bool `json:"success"`
	HasResult bool `json:"hasResult"`
	Result    any  `json:"result"`
}

func (i InvokeResult) BS() []byte {
	bs, _ := json.Marshal(i)
	return bs
}

func (i InvokeResult) HasValue() bool {
	return i.HasResult
}

func (i InvokeResult) Value() any {
	return i.Result
}
