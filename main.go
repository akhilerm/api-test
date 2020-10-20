package main

import (
	"fmt"
	calcv2 "github.com/akhilerm/api-test/v2/pkg/calc"
)

func main() {
	a := calcv2.NewVariables(10, 11, 12)
	fmt.Printf("Formatted variables = (%s) \n", a.GetFormattedString())
}
