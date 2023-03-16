package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello " + strconv.Itoa(i+1))
			wg.Done()
		}
	}()

	wg.Wait()
}
