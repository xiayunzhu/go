package main

import "fmt"

var (
	a = 1
	s = '1'
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Print(a, b, s)
}

func main() {
	fmt.Print("Heoolo wordl")
	variableInitialValue()
	variableZeroValue()
}
