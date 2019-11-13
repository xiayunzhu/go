package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "ccc",
		"cou":  "www",
		"site": "ddd",
	}
	m2 := make(map[string]int)

	var m3 map[string]int
   //遍历不保证顺序
	fmt.Println(m, m2, m3)
	for k, v := range m {
		fmt.Println(k, v)
	}

	//当key不存在的时候，获取value最初值
	if cou, ok := m["cou"]; ok {
		fmt.Println(cou)
	} else {
		fmt.Println("不存在")
	}

	//删除
	delete(m, "name")
	//判断key是否存在
	name, ok := m["name"]
	fmt.Println(name, ok)
	fmt.Println(len(m))
}
