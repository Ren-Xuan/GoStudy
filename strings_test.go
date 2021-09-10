package main

import (
	"fmt"
	"strconv"
	"strings"
)

func string_test() {
	var grade = 'A'
	fmt.Printf("%c", grade)

	fmt.Println(strings.Contains("abcd", "bc"))

	fmt.Println("fuck \n u")
	fmt.Println(`fuck \n u`) //原字符串
	fmt.Println(`abc
				 def
				 hij`)

	str := "Launch in T minus" + strconv.Itoa(10) + " seconds."
	//必须要这个函数才能转换 Integer to ASCII
	fmt.Println(str)
	//或者Atoi（ASCII to Integer)
	i, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("error atoi ", err)
	} else {
		fmt.Println(i)
	}
	//或者sprintf
	str2 := fmt.Sprintf("launch in T minus %v seconds.", 10)
	//用str就不行
	fmt.Println(str2)

}
