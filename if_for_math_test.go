package main

import (
	"fmt"
	"math"
	"math/rand"
)

func test() {
	fmt.Println("start")
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt8)
	if num := rand.Intn(3); num == 0 {
		fmt.Println("零")
	} else if num == 1 {
		fmt.Println("一")
	} else {
		fmt.Println("二")
	}

	for num := 0; num < 10; num++ {
		fmt.Println(num)
	}
}
