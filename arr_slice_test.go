package main

import "fmt"

func arr() bool {
	//go的拷贝是值拷贝
	//python拷贝是指针1
	values := [][3]int{}

	row1 := [...]int{1, 2, 3} //第一种声明并初始化数组方法
	var row2 [3]int           //第二种声明并初始化数组方法
	for i := 0; i < 3; i++ {
		row2[i] = i * i
	}
	fmt.Println("row1", row1)
	fmt.Println("row2", row2)
	values = append(values, row1)
	values = append(values, row2)
	//当slice的容量不足以append操作的时候，go会自动创建新数组并复制旧数组的内容
	fmt.Println(values[0])
	fmt.Println(values[0][0])

	arr := [...]string{ //不写...就是创建了一个slice
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
	}
	arr_a := arr[:4]
	arr_b := arr[4:6]
	arr_c := arr[6:]

	str := "abcde"
	substr := srt[1:2]

	//make语句
	dwarfs := make([]string, 0, 10)
	dump("dwarfs", dwarfs)
	return true
}
