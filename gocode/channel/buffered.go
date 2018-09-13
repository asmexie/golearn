//使用有缓冲通道和固定数目的goroutine来处理一堆工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  //要使用的goroutine数量
	taskLoad         = 10 //要处理的工作数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

//worker作为goroutine启动处理从有缓冲通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			//通道已经为空，并且已经被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		//显示开始工作
		fmt.Printf("Worker: %d: Started %s\n", worker, task)
		//随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		//显示任务完成
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
func main() {
	//创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	//启动goroutine来处理工作
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//当前所有工作都处理完，关闭通道以便goroutine退出
	close(tasks)
	wg.Wait()
}
