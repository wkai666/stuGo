package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x int64
	l sync.Mutex
	wg sync.WaitGroup
)

func add()  {
	x++
	wg.Done()
}

func mutexAdd()  {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

func atomicAdd()  {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main()  {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// go add()		// 普通版 add 并不是并发安全的
		// go mutexAdd()	// 加锁版 add，是并发安全的，但是加锁性能开销极大
		// go atomicAdd()	// 原子操作版 add ，并发安全的，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}