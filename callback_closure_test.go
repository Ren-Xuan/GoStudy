package main

import "fmt"

// 声明一个函数类型
type cb func(int) int
type cb2 func(int) float64

func callback_test() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
	//三种函数作为参数的情况
	var itof_func1 func(int) float64 = itof
	itof_func2 := itof
	var itof_func3 cb2 = itof

	fmt.Println(itof_func1(1))
	fmt.Println(itof_func2(1))
	fmt.Println(itof_func3(1))

	//闭包测试
	f(1)
	//闭包2
	f := func(i int) float64 {
		fmt.Println("closure test1")
		return float64(i)
	}
	f(2)
	//闭包3
	_ = func(i int) float64 {
		return float64(i)
	}(2)
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

func itof(i int) float64 {
	return float64(i)
}

//闭包
var f = func(i int) float64 {
	fmt.Println("closure test1")
	return float64(i)
}

//更奇怪的代码，会输出123321
func closureTest() {
	a := []int{1, 2, 3}
	for _, value := range a {
		fmt.Println(value)
		defer pppp(value)
	}
}
func pppp(value int) {
	fmt.Println(value)
}
