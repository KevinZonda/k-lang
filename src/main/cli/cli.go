package cli

import "github.com/alecthomas/kong"

type ParamModel struct {
	Repl struct{} `cmd:"" description:"Start REPL"`
	Ast  struct {
		Input  string `arg:"" description:"Input file" type:"path"`
		Output string `arg:"" description:"Output file" default:"" type:"path"`
	} `cmd:"" description:"Parse input file to AST"`
	Run struct {
		Input string `arg:"" type:"path"`
	} `cmd:""`
	Compile struct {
		Input  string `arg:"" type:"path"`
		Output string `arg:"" optional:"" type:"path"`
	} `cmd:""`
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
