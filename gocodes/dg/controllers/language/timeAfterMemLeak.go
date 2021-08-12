package language

// 用于测试time.After()的内存泄露, 因为每次调用均会产生一个chan, 若大量数据涌入会造成chan的内存泄露
// 正确方式: time.NewTimer()配合time.Reset()使用

import (
	"fmt"
	"sync"
	"time"
)

func timeAfterMemLeak() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			ch <- i
		}
	}()
LOOP:
	for {
		select {
		case now := <-time.After(3 * time.Minute):
			fmt.Println(now)
		case i := <-ch:
			//fmt.Print(i)
			if i == 1000000-1 {
				break LOOP
			}
		}
	}
	wg.Wait()
}
