package main

import (
	"fmt"
	"sync"
)

var x2 = 0

func incr2(group *sync.WaitGroup, ch chan bool) {
	ch <- true
	x2 += 1
	<-ch
	group.Done()
}
func main() {
	//var mutx sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incr2(&wg, ch)
	}
	wg.Wait()
	fmt.Println(x2)

}
