package main

import "fmt"

type celsius float64
type kelvin float64

const degree = 20

var temperature celsius = degree

//上面的作用就是提高可读性

//方法

func kelvinToCelsius(K kelvin) celsius {
	return celsius(K - 273.15)
}

func (k kelvin) Celsius() celsius {
	return celsius(k - 273.15)
}

func type_test() {
	var f func(k kelvin) celsius = kelvinToCelsius
	f(1)
	ktv := kelvinToCelsius
	fmt.Println(ktv(1))

	var k kelvin = 294.0
	var c celsius

	c = k.Celsius()

	fmt.Println(c)
}
