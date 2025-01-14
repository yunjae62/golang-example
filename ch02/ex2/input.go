package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input int

	n, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println("입력값:", input)
	}

	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int
	n2, err2 := fmt.Scanln(&a, &b)

	if err2 != nil {
		fmt.Println(err2)
		_, err := stdin.ReadString('\n')
		if err != nil {
			return
		}
	} else {
		fmt.Println("입력값:", a, b, n2)
	}
}
