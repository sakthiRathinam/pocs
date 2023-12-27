package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var max_num = 100000000
var num_of_prime_count int32 = 0
var current_num int32 = 0
var concurrency = 10

func main() {
	prime_count := execute_asynchronously_with_multiple_threads(max_num)
	// prime_count := prime_no_count_with_in_range(max_num)
	fmt.Printf("total no of prime number count: %d \n", prime_count)
}

func is_prime(num int) bool {
	if num < 2 {
		return false
	}
	num_is_prime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if (num % i) == 0 {
			num_is_prime = false
			break
		}
	}
	return num_is_prime
}

func do_work(wg *sync.WaitGroup, thread_num string) {
	defer wg.Done()
	start_time := time.Now()
	for {
		x := atomic.AddInt32(&current_num, 1)
		if x > int32(max_num) {
			break
		}
		if is_prime(int(x)) {
			atomic.AddInt32(&num_of_prime_count, 1)
		}
	}
	fmt.Printf("Thred %s took %s seconds\n", thread_num, time.Since(start_time))

}

func execute_asynchronously_with_multiple_threads(total_num int) int32 {
	start_time := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go do_work(&wg, strconv.Itoa(i))
	}
	wg.Wait()
	fmt.Printf("Function took %s seconds\n", time.Since(start_time))
	return num_of_prime_count
}
