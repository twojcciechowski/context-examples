package main

import (
	"context"
	"fmt"
	"time"
)

func withCancelExample() {
	fmt.Printf("------------------- Starting context withCancel example ------------------- \n")
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	start := time.Now()
	fmt.Printf("start: %v \n", start.Format("15:04:05"))

	go func() {
		f_cancel(ctx)
		cancelFunc()
	}()

	select {
	case <-time.After(4 * time.Second):
		{
			fmt.Printf("manual cancellation \n")
			cancelFunc()
		}
	}

	end := time.Now()
	fmt.Printf("end: %v \n", end.Format("15:04:05"))
	fmt.Printf("duration: %v \n", time.Duration(end.UnixNano()-start.UnixNano()).String())
	fmt.Printf("------------------- Ending context withCancel example ------------------- \n")
}

func f_cancel(ctx context.Context) {

	select {
	case <-ctx.Done():
		{
			fmt.Printf("f_cancel cancelled \n")
			return
		}
	case <-time.After(10 * time.Second):
		{
			fmt.Printf("function f_cancel finished\n")
			return
		}
	}

}
