package main

import "fmt"

func main() {
	s := "ffff赵娜"

	for _, b := range []byte(s) {
		fmt.Printf("%x  ", b)
	}
	//fmt.Printf("%X\n", []byte(s))
}
