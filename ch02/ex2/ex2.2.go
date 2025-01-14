package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var b int16
	var c float64 = 3.14
	d := 5
	d++

	fmt.Println(b, c, d)

	d = int(c)
	fmt.Println(d)
}
