package goroutines

import (
	"fmt"
	"math"
	"sync"
)

// func isPrime(start, end int, channel chan int, done chan bool) {
// 	for start < end {
// 		isPrime := true
// 		limit := int(math.Sqrt(float64(start)))
// 		for i := 2; i <= limit; i++ {
// 			if start%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			channel <- start
// 		}
// 		start++
// 	}
//
// 	done <- true
// }

// Better way
func isPrimeUsingSync(start, end int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for start < end {
		isPrime := true
		limit := int(math.Sqrt(float64(start)))
		for i := 2; i <= limit; i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			channel <- start
		}
		start++
	}
}

func GoroutinesExample() {
	channel := make(chan int)
	// done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(3)

	// go isPrime(2, 30, channel, done)
	// go isPrime(30, 70, channel, done)
	// go isPrime(70, 101, channel, done)
	go isPrimeUsingSync(2, 30, channel, &wg)
	go isPrimeUsingSync(30, 70, channel, &wg)
	go isPrimeUsingSync(70, 101, channel, &wg)

	// IMPORTANT: the below lines <-done is waiting that means the line is waiting for value to be
	// returned similar to "await"
	// go func() {
	// 	<-done
	// 	<-done
	// 	<-done
	// close(channel)
	// }()

	go func() {
		wg.Wait()
		close(channel)
	}()

	fmt.Println("The prime number found in parallel using go goroutines and channels are - ")
	for val := range channel {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
