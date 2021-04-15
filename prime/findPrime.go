package main

import (
	"fmt"
	"math"
)

func FindPrime() (result []int) {
	for i := 1; i <= 200000; i++ {
		result = append(result, i)
		for j := 2; j < i; j++ {
			if i%j == 0 {
				result = result[:len(result)-1]
				break
			}
		}
	}
	return
}

func FindPrime1() (result []int) {
	for i := 1; i <= 200000; i++ {
		result = append(result, i)
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				result = result[:len(result)-1]
				break
			}
		}
	}
	return
}

func FindPrime2() (result []int) {
	primeChan := make(chan int, 10)
	exitChan := make(chan interface{}, 10)
	for i := 1; i <= 200000; i++ {
		go func(num int, c chan int) {
			var j int
			for j = 2; j < num; j++ {
				if num%j == 0 {
					break
				}
			}
			if j >= num {
				c <- num
			}
			exitChan <- nil
		}(i, primeChan)
	}
	go func() {
		for i := 0; i < 200000; i++ {
			<-exitChan
		}
		close(exitChan)
		close(primeChan)
	}()
	for {
		if num, ok := <-primeChan; ok {
			result = append(result, num)
		} else {
			break
		}
	}
	return
}

func FindPrime3() (result []int) {
	result = append(result, 1)
	origin, wait := make(chan int), make(chan struct{})
	findIt(origin, wait, 1, &result)
	for i := 2; i < 200000; i++ {
		origin <- i
	}
	close(origin)
	<-wait
	return
}

func findIt(seq <-chan int, wait chan struct{}, level int, primes *[]int) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		//fmt.Println("[", level, "]: ", prime)
		*primes = append(*primes, prime)
		out := make(chan int)
		findIt(out, wait, level+1, primes)

		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}

// shared state
func FindPrime4() (result []int) {
	start, end := 0, 200000 // [1,200000]
	times, length := 7, end-start
	step := length / times
	for i := 0; i < times; i++ {
		go isPrimes(start+1, start+step, &result)
		start = start + step
	}
	go isPrimes(start+1, end, &result)
	return
}

func isPrimes(start, end int, outcome *[]int) {
	for i := start; i <= end; i++ {
		*outcome = append(*outcome, i)
		for j := 2; j < i; j++ {
			if i%j == 0 {
				*outcome = (*outcome)[:len(*outcome)-1]
				break
			}
		}
	}
}

func FindPrime5() (result []int) {
	start, end := 0, 200000 // [1,200000]
	times, length := 20, end-start
	step := length / times
	outcome := make(chan []int, times)
	for i := 0; i < times; i++ {
		tmp := []int{}
		go isPrimes1(start+1, start+step, outcome)
		start = start + step
		result = append(result, tmp...)
	}
	go isPrimes1(start+1, end, outcome)
	for i := 0; i <= times; i++ {
		result = append(result, <-outcome...)
	}
	return
}
func isPrimes1(start, end int, writeto chan []int) {
	outcome := []int{}
	for i := start; i <= end; i++ {
		outcome = append(outcome, i)
		for j := 2; j < i; j++ {
			if i%j == 0 {
				outcome = (outcome)[:len(outcome)-1]
				break
			}
		}
	}
	writeto <- outcome
}

func main() {
	fmt.Println(len(FindPrime5()))
	/*
		fmt.Println(len(FindPrime3()))
		fmt.Println(len(FindPrime2()))
		answer := FindPrime1()
		fmt.Println(len(answer))
		fmt.Println(FindPrime())
	*/
}
