package goroutines

import (
	"fmt"
	"time"
)

func SelectExample() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from channel 1"
		close(channel1)
	}()

	// Goroutine 2
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from channel 2"
		close(channel2)
	}()

	defer func() {
		fmt.Println("Select finished waiting!")
	}()

	// Select will wait for the first available channel
	i := 0
	for {
		i++
		// fmt.Println("Loop number", i)
		if channel1 == nil && channel2 == nil {
			fmt.Println("exit after loop number", i)
			return
		}

		select {
		case msg1, ok := <-channel1:
			if !ok {
				channel1 = nil
				continue
			}
			fmt.Println("Loop", i, "got:", msg1)
		case msg2, ok := <-channel2:
			if !ok {
				channel2 = nil
				continue
			}
			fmt.Println("Loop", i, "got:", msg2)
		}
	}
}

