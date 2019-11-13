package main

import "fmt"

//切片
//slice本身没有值，对原数组的view

func updateSlice(s []int) {

	//值改变，会影响底层数组的变化
	s[0] = 100
	fmt.Println(s)
}

//向slice添加元素
//超出底层数组的cap，会拷贝原来的数组，成为一个新的数组
func add(s []int) {
	s3 := append(s, 10)
	fmt.Println(s3)
	s4 := append(s3, 11)
	fmt.Println(s4)
	s5 := append(s4, 11)
	fmt.Println(s5)
}
func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[:])
	//s1:=arr[2:]
	s2 := arr[:]
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:8] //超出当前slice长度，slice只要不超过原数组的长度可以向后扩展，不能向前扩展
	fmt.Println(s2)
	//updateSlice(s1)
	updateSlice(s2)
	add(s2)
}
