package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)

	go func() {
		for i := 5; i >= 0; i-- {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for data := range ch {
		fmt.Println("data is: ", data)
		if data == 2 {
			break
		}
	}

	data2 := <- ch
	fmt.Println("data2 is: ", data2)
	fmt.Println("all done")
}
