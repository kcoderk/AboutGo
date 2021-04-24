package main

import (
	"fmt"
	"sync"
)

var x = 0

func incr(group *sync.WaitGroup) {

	x += 1
	group.Done()
}
func main() {
	//var mutx sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incr(&wg)
	}
	wg.Wait()
	fmt.Println(x)

}
