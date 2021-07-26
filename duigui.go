package main

import "fmt"

func main()  {
	digui(10)
}

func digui(n int) {
	if n == 0 {
		return
	} else {
		fmt.Println(n)
		digui(n-1)
	}
}
