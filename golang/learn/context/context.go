package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	test2()
}
func test1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		fmt.Println("in")

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

	}()
	time.Sleep(time.Second * 2)
}

func test2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go func() {
		fmt.Println("in")

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	}()
	time.Sleep(time.Second * 3)
}
