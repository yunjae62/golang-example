package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func ChangeData(arg Student) {
	arg.Name = "바뀐 이름"
}

func ChangeDataPointer(arg *Student) {
	arg.Name = "바뀐 이름"
}

func main() {
	var a Student
	a.Name = "철수"
	a.Age = 16

	b := Student{"영희", 17}

	c := Student{Name: "민수"}
	c.Age = 18

	fmt.Println(a)
	fmt.Println(b.Name)
	fmt.Println(c)

	var p1 *Student = &a

	fmt.Printf("%p %v\n", p1, *p1)

	ChangeData(a)
	fmt.Println(a)

	ChangeDataPointer(&a)
	fmt.Println(a)

	var d *Student = new(Student)
	fmt.Println(*d)

	var e = &Student{}
	fmt.Println(*e)
}
