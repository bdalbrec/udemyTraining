// can still lead to race conditions and misordered output, but the final result will be correct
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		//mutex.Lock(
		//counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter) // this causes the race condition as it is not thread specific withouth the mutex
		//mutex.Unlock()
	}
	wg.Done()
}
