package main

import "fmt"

func main()  {
	var index, result int
	for i := 0; i < 10; i++ {
		index, result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", index, result)
	}
}

func fibonacci(a int) (n int, res int) {
	if a <= 1 {
		res = 1
	} else {
		_, res1 := fibonacci(a - 1)
		_, res2 := fibonacci(a - 2)
		res = res1 + res2
	}
	return a, res
}
