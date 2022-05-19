package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	prob := flag.String("problem", "Problem1", "problem to run")
	arg := flag.Int("arg", 1, "argument to pass to problem")
	flag.Parse()

	m := map[string]func(int) int{
		"Problem1": problem1,
		"Problem2": problem2,
		"Problem3": problem3,
	}

	if m[*prob] == nil {
		fmt.Println("You forgot to add the problem to the map")
		os.Exit(0)
	}

	start := time.Now()
	fmt.Printf("%d (%v)", run(m[*prob], *arg), time.Since(start))
}

func run(f func(int) int, a int) int {
	return f(a)
}

func problem1(num int) int {
	sum := 0
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func problem2(limit int) int {
	num1 := 0
	num2 := 1
	var fib int
	sum := 0
	for fib <= limit {
		fib = num1 + num2
		num1 = num2
		num2 = fib
		if fib%2 == 0 {
			sum += fib
		}
	}
	return sum
}

func problem3(num int) int {
	var lastPrime int
	for prime := 2; prime <= num; prime++ {
		if !primeCheck(prime) {
			continue
		}
		for num%prime == 0 {
			num /= prime
		}
		lastPrime = prime
	}
	return lastPrime
}

func primeCheck(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
