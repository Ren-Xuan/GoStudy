package main

import "fmt"

tye report struct{
	sol int
	temperature temperature
	location location

}

type temperature struct{
	high,low celsius
}

type location struct{
	lat,long float64
}

type celsius float64

func struct_test(){

}