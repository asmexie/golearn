//使用无缓冲通道模拟4个goroutine接力比赛
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Runner(baton chan int) {
	var newRuner int
	//等待接力棒
	runner := <-baton
	//开始绕跑到跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	//创建下一位跑步者
	if runner != 4 {
		newRuner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRuner)
		go Runner(baton)
	}
	//围绕跑道跑
	time.Sleep(100 * time.Millisecond)
	//比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}
	//将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRuner)
	baton <- newRuner
}
func main() {
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()
}
