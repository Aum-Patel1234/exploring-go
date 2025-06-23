package main

import (
	"fmt"
	"goroutines/goroutines"
)

func main() {
	fmt.Println()
	goroutines.GoroutinesExample()

	goroutines.SelectExample()

	// NEXT TIME - implement fibonacci sequence using goroutines like for n=4,5,2,78 concurrently
}

// NOTE: Discover more about sync.Mutex and sync.RWMutex
// func reader(ch <-chan int) { // Read-only channel
// }
//
// func writer(ch chan<- int) { // Write-only channel
// }
