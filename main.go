package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Operation one
	operationOneCtxDeadline := time.Now().Add(5 * time.Second)
	operationOneCtx, operationOneCtxCancel := context.WithDeadline(context.Background(), operationOneCtxDeadline)
	operationOneCtx = context.WithValue(operationOneCtx, "operation_type", "one")
	defer operationOneCtxCancel()
	go operationOne(operationOneCtx)

	// Operation two
	operationTwoCtxDeadline := time.Now().Add(10 * time.Second)
	operationTwoCtx, operationTwoCtxCancel := context.WithDeadline(context.Background(), operationTwoCtxDeadline)
	operationTwoCtx = context.WithValue(operationTwoCtx, "operation_type", "two")
	defer operationTwoCtxCancel()
	go operationTwo(operationTwoCtx)

	time.Sleep(20 * time.Second)
}

func operationOne(ctx context.Context) {
	counter := 1
	operation_type := ctx.Value("operation_type")

	operationOneChildCtx := context.WithValue(ctx, "operation_type", "one child")
	go operationOneChild(operationOneChildCtx)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("operation %s context canceled\n", operation_type)
			return
		default:
			fmt.Printf("operation %s counter: %d\n", operation_type, counter)
			time.Sleep(500 * time.Millisecond)
			counter++
		}
	}
}

func operationOneChild(ctx context.Context) {
	operation_type := ctx.Value("operation_type")

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("context canceled for %s\n", operation_type)
			return
		default:
			fmt.Println("child of operation one")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func operationTwo(ctx context.Context) {
	counter := 1
	operation_type := ctx.Value("operation_type")

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("operation %s context canceled\n", operation_type)
			return
		default:
			fmt.Printf("operation %s counter: %d\n", operation_type, counter)
			time.Sleep(250 * time.Millisecond)
			counter++
		}
	}
}
