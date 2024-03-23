package httpExternelServer

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel/httpExternel"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"io"
	"net/http"
	"reflect"
)

func NewFuncPack(name string) *FuncPack {
	return &FuncPack{packName: name}
}

type FuncPack struct {
	fxs      map[string]any
	fxa      map[string]externel.FuncArgs
	packName string
}

func (p *FuncPack) AppendFx(fx any) *FuncPack {
	return p.AppendFxWithName(FuncName(fx), fx)
}

func (p *FuncPack) AppendFxWithName(name string, fx any) *FuncPack {
	if p.fxs == nil {
		p.fxs = map[string]any{}
	}
	if p.fxa == nil {
		p.fxa = map[string]externel.FuncArgs{}
	}
	p.fxs[name] = fx
	p.fxa[name] = readFuncArg(fx)
	return p
}

func (p *FuncPack) StartServer(listen string) {
	h := http.NewServeMux()
	for k, fx := range p.fxs {
		if fx == nil {
			return
		}
		hdlr := createFuncHandler(reflect.ValueOf(fx), p.fxa[k])
		if hdlr == nil {
			continue
		}
		h.HandleFunc("POST /call/"+p.packName+"/"+k, hdlr)
	}
	h.HandleFunc("GET /"+p.packName, func(w http.ResponseWriter, r *http.Request) {
		rsp := map[string]externel.PackageInfoElement{}
		for name, _ := range p.fxs {
			rsp[name] = externel.PackageInfoElement{
				Type: externel.Function,
				Args: p.fxa[name],
			}
		}

		bs, _ := json.Marshal(rsp)
		w.WriteHeader(200)
		w.Write(bs)
	})
	http.ListenAndServe(listen, h)
}

func createFuncHandler(f reflect.Value, fxArg externel.FuncArgs) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req httpExternel.FuncCallRequest
		err := deserialise(r.Body, &req)
		panicx.PanicIfNotNil(err, err)
		req.Args = externel.ConvertArg(req.Args, fxArg)
		var args []reflect.Value
		for _, v := range req.Args {
			args = append(args, reflect.ValueOf(v))
		}

		vals := f.Call(args)
		resp := externel.InvokeResult{
			Success:   true,
			HasResult: len(vals) != 0,
		}

		if resp.HasResult {
			var rspVal []any
			for _, v := range vals {
				rspVal = append(rspVal, v.Interface())
			}
			if len(rspVal) == 1 {
				resp.Result = rspVal[0]
			} else {
				resp.Result = rspVal
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp.BS())
	}

}

func deserialise(r io.Reader, v any) error {
	bs, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}
