package httpExternel

import (
	"bytes"
	"encoding/json"
	"io"
)

type FuncCallRequest struct {
	FuncName string
	Args     []any
}

func (f FuncCallRequest) IO() io.Reader {
	bs, e := json.Marshal(f)
	if e != nil {
		panic(e)
	}
	return bytes.NewReader(bs)
}
