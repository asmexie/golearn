//使用互斥锁，来定义一段需要同步访问的代码临界区资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//同一时刻只允许一个goroutine进入临界区
		mutex.Lock()
		{
			value := counter
			//当前goroutine线程退出，并放回队列
			runtime.Gosched()
			value++
			counter = value
		}
		//释放锁，允许其他正在等待的goroutine进入临界区
		mutex.Unlock()
	}
}
func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter:%d\n", counter)
}
