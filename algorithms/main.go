package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	result := findSum(99999999)
	timeEnd := time.Now()
	fmt.Println("O resultado", result)
	fmt.Println(timeEnd.Sub(timeStart).Milliseconds())
}

func findSum(n int) int {
	return n * (n + 1) / 2
}

func findSumIterate(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}
