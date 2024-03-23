package jout

import (
	"encoding/json"
	"fmt"
)

func Println(items ...any) {
	for _, item := range items {
		bs, _ := json.MarshalIndent(item, "", "    ")
		fmt.Println(string(bs))
	}
}

func PrintAndRet[T any](item T) T {
	Println(item)
	return item

}
