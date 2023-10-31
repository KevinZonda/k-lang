package cli

import "github.com/alecthomas/kong"

type ParamModel struct {
	Repl struct{} `cmd:"" description:"Start REPL"`
	Ast  struct {
		Input  string `arg:"" description:"Input file"`
		Output string `arg:"" description:"Output file"`
	} `cmd:"" description:"Parse input file to AST"`
}

var param *ParamModel
var ctx *kong.Context

func ParseParam() {
	param = new(ParamModel)
	ctx = kong.Parse(param)
}

func Ctx() *kong.Context {
	return ctx
}

func Param() *ParamModel {
	return param
}
