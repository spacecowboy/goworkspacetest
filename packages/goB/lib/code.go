package lib

import libA "example.com/packages/goA/lib"

func FuncB() string {
	return "v2.FuncB feature fix"
}

type StructB struct {
	FieldB libA.StructA
}
