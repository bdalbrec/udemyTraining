package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// init is a special function to set up a program before execution
// this uses all the CPU cores, after Go 1.5 this is default and not needed. Before 1.5 it only used 1 core.
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}
