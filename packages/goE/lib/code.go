package lib

import libA "example.com/packages/goA/lib"

func FuncE() string {
	return "FuncE updated"
}

type StructE struct {
	FieldE libA.StructA
}
