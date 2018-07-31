//使用无缓冲通道模拟网球比赛
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		//等待球被击打过来
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		//选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//关闭通道表示我们输了
			close(court)
			return
		}
		//显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		//将球打给对手
		court <- ball
	}
}
func main() {
	//创建一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)
	court <- 1
	wg.Wait()

}
