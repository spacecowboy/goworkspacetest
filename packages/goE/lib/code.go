package lib

import libA "example.com/packages/goA/lib"

func FuncE() string {
	return "FuncE updated fix"
}

type StructE struct {
	FieldE libA.StructA
}
