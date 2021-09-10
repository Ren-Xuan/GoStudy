package main

import "fmt"

//map是引用传递
func map_test() {
	temperature := map[sting]int{
		"earth": 15,
		"mars":  -65,
	}
	temperature["venus"] = 465
	fmt.Println(temperature["moon"])
	temperatureII := temperature
	delete(temperatureII, "earth")
}
