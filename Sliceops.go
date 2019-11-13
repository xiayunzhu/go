package main

import "fmt"

func printSlice(s []int) {

	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
func main() {
	var s []int //zero value for silce is nil  不赋值的时候 是空[]

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	//新建长度为16的slice
	s2 := make([]int, 16)
	printSlice(s2)
	//新建长度为10的slice，cap为32的slice
	s3 := make([]int, 10, 32)
	printSlice(s3)
	//拷贝数组
	copy(s2,s1)
	fmt.Println(s2)
	//删除一些元素
	s2=append(s2[:3],s2[4:]...)
	fmt.Println(s2)
	printSlice(s2)

	//
	

}
