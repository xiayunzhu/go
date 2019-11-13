package main

import "fmt"

//数组是值类型，长度不一样，为不同类型，例如arr [5]int和arr [3]int为不同类型
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//指针传递
func printArra(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//arr作为参数，会拷贝数组
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{2, 4, 6, 8, 10} //让系统来判断数组长度
	var grid [4][5]int               //4行5列的数组

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//数组遍历

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//range
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	//只想要值
	for _, v := range arr3 {
		fmt.Println(v)
	}
	printArray(arr3)
	printArra(&arr3)
}
