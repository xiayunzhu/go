package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int)string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong  score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}
func main() {

	fmt.Print(grade(100))
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Print(string(contents))
	} else {
		fmt.Print("cannot print file contents:", err)
	}
}
