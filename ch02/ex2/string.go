package main

import (
	"fmt"
	"strings"
)

func main() {
	// 고랭은 UTF-8 사용

	str1 := `안녕하세요.\n 백틱쓰면 \n작동안함
근데 줄바꿈은 됨
ㅋ`
	fmt.Println(str1)
	fmt.Println(len(str1))

	var ch1 int32 = '한'
	var ch2 rune = '두'

	fmt.Printf("%c %c\n", ch1, ch2)
	fmt.Println(ch1)
	fmt.Println(ch2)

	str2 := "Hello"
	runes2 := []rune{'H', 'e', 'l', 'l', 'o'}

	fmt.Println(str2)
	fmt.Println(string(runes2))

	runes3 := []rune(str2)

	for i := 0; i < len(runes3); i++ {
		fmt.Printf("%c\n", runes3[i])
	}

	for _, v := range str2 {
		fmt.Printf("%c\n", v)
	}

	var builder strings.Builder
	for _, ch := range str2 {
		builder.WriteRune(ch)
	}

	fmt.Println(builder.String())
}
