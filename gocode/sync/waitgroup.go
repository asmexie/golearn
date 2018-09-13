package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	timeout := make(chan bool)
	data := make(chan int, 1)
	ticker := time.NewTicker(time.Second * 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	i := 0
	go func(wg1 *sync.WaitGroup, i int) {
		for {
			select {
			case <-timeout:
				log.Println("timeout, close goroutine")
				wg1.Done()
				return
			case <-ticker.C:
				log.Println("write data", i)
				data <- i
				log.Println("write data success ", i)
				i++
			}
		}
	}(wg, i)

	after := time.After(time.Second * 5)
	wg.Add(1)
	go func(wg1 *sync.WaitGroup) {
		for {
			select {
			case x := <-data:
				log.Println("read from data ", x)
			case <-after:
				log.Println("write timeout")
				timeout <- true
				log.Println("write timeout success")
				close(data)
				wg1.Done()
				return
			}
		}
	}(wg)
	wg.Wait()
	time.Sleep(time.Second * 5)
	log.Println("finish")
}
