package main

import "fmt"

func main()  {
	n := 6
	fmt.Printf("%d jicheng result is %d\n", n, ji(n))
}

func ji(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * ji( n-1 )
	}
}
