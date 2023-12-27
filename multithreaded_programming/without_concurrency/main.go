package main

import (
	"fmt"
	"math"
	"time"
)

var max_num = 100000000
var num_of_prime_count = 0

func main() {
	prime_count := prime_no_count_with_in_range(max_num)
	fmt.Printf("total no of prime number count: %d \n", prime_count)
}

func is_prime(num int) bool {
	num_is_prime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if (num % i) == 0 {
			num_is_prime = false
			break
		}
	}
	return num_is_prime
}

func prime_no_count_with_in_range(num int) int {
	start_time := time.Now()
	for i := 2; i <= num; i++ {
		if is_prime(i) {
			num_of_prime_count += 1
		}
	}
	end_time := time.Now()
	duration := end_time.Sub(start_time)
	fmt.Printf("Function took %f seconds\n", duration.Seconds())
	return num_of_prime_count
}
