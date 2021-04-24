package main

import (
	"fmt"
	"sync"
)

var x1 = 0

func incr1(group *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x1 += 1
	mutex.Unlock()
	group.Done()
}
func main() {
	//var mutx sync.Mutex
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incr1(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println(x1)

}
