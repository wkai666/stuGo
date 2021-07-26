package main

import "fmt"

func main()  {

	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		fmt.Println("end go routine")
	}()

	fmt.Println("wait goroutine")

	data := <- ch
	fmt.Println("data is: ", data)
	fmt.Println("all done")
}
