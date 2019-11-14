package main

import "fmt"

type Retiever interface {
	Get(url string) string
}

//接口由使用者定义
func download(r Retiever) string {
	return r.Get("WWWWWWW")
}
func main() {

	var r Retiever
	fmt.Println(download(r))
}
