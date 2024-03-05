package module

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/hjson/hjson-go/v4"
)

func LoadFromText(text string) Mod {
	var m Mod
	err := hjson.Unmarshal([]byte(text), &m)
	if err != nil {
		panic(err)
	}
	return m
}

func LoadAsMod(text string) (mod Mod, ok bool) {
	err := hjson.Unmarshal([]byte(text), &mod)
	ok = err == nil
	return
}

func LoadFromPath(path string) Mod {
	bs, err := iox.ReadAllText(path)
	if err != nil {
		panic(err)
	}
	return LoadFromText(bs)
}
