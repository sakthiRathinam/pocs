package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var max_num = 100000000
var num_of_prime_count int32 = 0

func main() {
	prime_count := execute_in_batch_wise(max_num)
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

func do_batch(wg *sync.WaitGroup, start_range int, end_range int) int32 {
	defer wg.Done()
	start_time := time.Now()
	for i := start_range; i < end_range; i++ {
		if is_prime(i) {
			atomic.AddInt32(&num_of_prime_count, 1)
		}
	}
	end_time := time.Now()
	duration := end_time.Sub(start_time)
	fmt.Printf("%d to %d batch took %f seconds\n", start_range, end_range, duration.Seconds())
	return num_of_prime_count
}

func execute_in_batch_wise(total_num int) int32 {
	var batch_size = int(100000000 / 10)
	start_time := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		start_range := batch_size * i
		end_range := (batch_size * i) + batch_size
		go do_batch(&wg, start_range, end_range)
	}
	wg.Wait()
	end_time := time.Now()
	duration := end_time.Sub(start_time)
	fmt.Printf("Function took %f seconds\n", duration.Seconds())
	return num_of_prime_count
}
