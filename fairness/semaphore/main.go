package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000
var CONCURRENCY = 10
var totalPrimeNumbers int32 = 0
var currentNum int32 = 2
var semaphore int = 0

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&totalPrimeNumbers, 1)
}

func main() {
	start := time.Now()
	semaphore := make(chan int, 10)
	for i := 3; i < MAX_INT; i++ {
		semaphore <- 1
		num := atomic.AddInt32(&currentNum, 1)
		go checkPrime(int(num))
		<-semaphore
	}
	fmt.Println("total number of prime  numbers are ", totalPrimeNumbers)
	fmt.Println("totla time is : ", time.Since(start))

}
