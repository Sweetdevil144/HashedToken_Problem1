package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time taken to execute main function:", time.Since(start))
	}()
	arr := []int{1, 2, 3, 4, 5}
	for _, element := range arr {
		go findOdd(element)
	}

	time.Sleep(time.Second) // Wait for goroutines to finish
}

func findOdd(num int) {
	fmt.Println("Finding odd for", num, ":", num%2 != 0)
}
