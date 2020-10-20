package main

import (
	"fmt"
	"github.com/akhilerm/api-test/pkg/calc"
)

func main() {
	a := calc.NewVariables(10, 11)
	fmt.Printf("Formatted variables = (%s) \n", a.GetFormattedString())
}
