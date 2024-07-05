package httpExternalServer

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/common"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternal"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"reflect"
)

func NewFuncPack(name string) *FuncPack {
	return &FuncPack{packName: name}
}

type FuncPack struct {
	fxs      map[string]any
	fxa      map[string]common.FuncArgs
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
		p.fxa = map[string]common.FuncArgs{}
	}
	p.fxs[name] = fx
	p.fxa[name] = readFuncArg(fx)
	return p
}

func (p *FuncPack) StartServer(listen string) {
	h := gin.New()
	for k, fx := range p.fxs {
		if fx == nil {
			return
		}
		hdlr := createFuncHandler(reflect.ValueOf(fx), p.fxa[k])
		if hdlr == nil {
			continue
		}
		h.POST("/"+p.packName+"/"+k, hdlr)
	}
	h.GET("/"+p.packName, func(ctx *gin.Context) {
		rsp := map[string]common.PackageInfoElement{}
		for name, _ := range p.fxs {
			rsp[name] = common.PackageInfoElement{
				Type: common.Function,
				Args: p.fxa[name],
			}
		}
		ctx.JSON(200, rsp)

	})
	h.Run(listen)
}

func createFuncHandler(f reflect.Value, fxArg common.FuncArgs) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req httpExternal.FuncCallRequest
		err := ctx.BindJSON(&req)
		panicx.NotNilErr(err)

		req.Args = common.ConvertArg(req.Args, fxArg)
		var args []reflect.Value
		for _, v := range req.Args {
			args = append(args, reflect.ValueOf(v))
		}

		vals := f.Call(args)
		resp := common.InvokeResult{
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
		ctx.JSON(http.StatusOK, resp)
	}

}

func deserialise(r io.Reader, v any) error {
	bs, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}
