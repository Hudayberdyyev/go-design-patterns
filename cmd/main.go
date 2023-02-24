package main

import (
	"fmt"
	"go-design-patterns/creational/singleton"
)

func main() {
	counter := singleton.GetInstance()
	counter.AddOne()
	fmt.Println(counter.GetCount())
}
