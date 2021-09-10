package main

import (
	"fmt"
)

type Sleeper interface {
	Sleep()
}
type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("dog %s is sleeping\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("cat %s is sleeping\n", c.Name)
}

func AnimalSleep(s Sleeper) {
	s.Sleep()
}
func main2() {
	fmt.Printf("Fuck you\n")

	var a float32 = 1.1
	b := int(a)

	fmt.Printf("new b =%d\n", b)

	var s Sleeper
	dog := Dog{Name: "mang ge"}
	cat := Cat{Name: "Tom"}
	s = dog
	AnimalSleep(s)
	s = cat
	AnimalSleep(s)

}
