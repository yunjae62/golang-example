package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
