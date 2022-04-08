package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mt sync.Mutex
	h := map[int]int{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			mt.Lock()
			h[0] = i

			mt.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(h)
}
