package main

import (
	"fmt"
	"time"
)

func main()  {
	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	// 还未开始就已经取消了
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
}
