package main

import "fmt"

type User struct {
	id   int
	name string
}

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	fmt.Println(c)
}

func main() {
	//u := User{1, "Tom"}
	//var i interface{} = u //相当于持有了u 的动态值，但是i 的动态类型还是interface
	//var i interface{}
	//fmt.Println(i)
	//fmt.Printf("%T", i)

	test(red)
	test(100)
	test(1)
	//a := 1
	//test(a)
}
