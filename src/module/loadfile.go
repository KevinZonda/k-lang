package module

import "github.com/hjson/hjson-go/v4"

func LoadFromText(text string) Mod {
	var m Mod
	err := hjson.Unmarshal([]byte(text), &m)
	if err != nil {
		panic(err)
	}
	return m
}
