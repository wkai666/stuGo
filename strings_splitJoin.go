package main

import (
	"fmt"
	"strings"
)

func main()  {

	str := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)
	fmt.Printf("Splitted in slice, type is %T, value is %v\n", s1, s1)
	for _, val := range s1 {
		fmt.Printf("%s\n", val)
	}

	fmt.Println()

	str2 := "GO|The ABC of Go |25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("str2 is: %s\n", str2)
	fmt.Printf("The sl2 type is %T, value is %v\n", sl2, sl2)
	for _, val := range sl2 {
		fmt.Printf("%s\n", val)
	}
	fmt.Println()

	str3 := strings.Join(sl2, ";")
	fmt.Printf("str3 is : %s\n", str3)
}