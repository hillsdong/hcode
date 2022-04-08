package main

import (
	"fmt"
	"time"
)

func main() {
	a := map[int64]string{}
	for i := 0; i < 400000; i++ {
		a[2021032415125227083+int64(i)] = "#FFFFFF"
	}
	fmt.Println("done")
	time.Sleep(time.Duration(2000) * time.Second)

}
