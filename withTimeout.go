package main

import (
	"context"
	"fmt"
	"time"
)

func withTimeoutExample() {
	fmt.Printf("------------------- Starting context withTimeout example ------------------- \n")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(9*time.Second))
	defer cancelFunc()

	start := time.Now()
	fmt.Printf("start: %v \n", start.Format("15:04:05"))

	f_timeout(ctx)

	end := time.Now()
	fmt.Printf("end: %v \n", end.Format("15:04:05"))
	fmt.Printf("duration: %v \n", time.Duration(end.UnixNano()-start.UnixNano()).String())
	fmt.Printf("------------------- Ending context withTimeout example ------------------- \n")
}

func f_timeout(ctx context.Context) {

	select {
	case <-ctx.Done():
		{
			fmt.Printf("f_timeout cancelled or timedout\n")
			return
		}
	case <-time.After(10 * time.Second):
		{
			fmt.Printf("function f_timeout finished\n")
			return
		}
	}

}
