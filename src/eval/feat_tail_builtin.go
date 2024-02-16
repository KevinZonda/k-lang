package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"strconv"
	"strings"
)

func (e *Eval) builtInString(recv string, fc node.FuncCall) any {
	args := e.evalExprs(fc.Args...)
	switch fc.Caller.Value {
	case "replace":
		return strings.ReplaceAll(recv, args[0].(string), args[1].(string))
	case "trim":
		return strings.TrimSpace(recv)
	case "upper":
		return strings.ToUpper(recv)
	case "lower":
		return strings.ToLower(recv)
	case "split":
		return strings.Split(recv, args[0].(string))
	case "contains":
		return strings.Contains(recv, args[0].(string))
	case "starts_with":
		return strings.HasPrefix(recv, args[0].(string))
	case "ends_with":
		return strings.HasSuffix(recv, args[0].(string))
	case "to_int":
		i, err := strconv.Atoi(recv)
		if err != nil {
			panic("not valid int")
		} else {
			return i
		}
	case "to_num":
		f, err := strconv.ParseFloat(recv, 64)
		if err != nil {
			panic("not valid number")
		}
		return f
	case "to_bool":
		b, err := strconv.ParseBool(recv)
		if err != nil {
			panic("not valid bool")
		}
		return b
	case "len":
		return len([]rune(recv))
	}
	panic("Unknown built-in string func: " + fc.Caller.Value)
}

func (e *Eval) builtInArray(recv []any, fc node.FuncCall) any {
	args := e.evalExprs(fc.Args...)
	switch fc.Caller.Value {
	case "append":
		return append(recv, args[0])
	case "pop":
		if len(recv) == 0 {
			return []any{recv, nil}
		}
		fst := recv[0]
		return []any{recv[1:], fst}
	case "join":
		sep := args[0].(string)
		sb := strings.Builder{}
		for i, v := range recv {
			if i != 0 {
				sb.WriteString(sep)
			}
			sb.WriteString(fmt.Sprint(v))
		}
		return sb.String()
	case "len":
		return len(recv)
	}
	panic("Unknown built-in array func: " + fc.Caller.Value)
}

func (e *Eval) builtInMap(recv map[any]any, fc node.FuncCall) any {
	args := e.evalExprs(fc.Args...)
	switch fc.Caller.Value {
	case "keys":
		keys := make([]any, 0, len(recv))
		for k := range recv {
			keys = append(keys, k)
		}
		return keys
	case "values":
		values := make([]any, 0, len(recv))
		for _, v := range recv {
			values = append(values, v)
		}
		return values
	case "remove":
		delete(recv, args[0])
		return nil
	case "len":
		return len(recv)
	}
	panic("Unknown built-in map func: " + fc.Caller.Value)
}
