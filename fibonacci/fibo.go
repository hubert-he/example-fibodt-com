package main

import "fmt"

//var cache map[int]int
var cache map[int]int = make(map[int]int)

func main() {
	fmt.Println(Fib1(6), Fib2(6))
}

func Fib1(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return Fib1(num-1) + Fib1(num-2)
}

func Fib2(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	if _, ok := cache[num-1]; !ok {
		cache[num-1] = Fib2(num - 1)
	}
	if _, ok := cache[num-2]; !ok {
		cache[num-2] = Fib2(num - 2)
	}
	return cache[num-1] + cache[num-2]
}

func Fib3(num int) (result int) {
	switch {
	case num < 0:
		return -1
	case num == 0:
		return 0
	case num <= 2:
		return 1
	default:
		i, j := 1, 1
		for k := 3; k <= num; k++ {
			result = i + j
			i, j = j, result
		}
	}
	return
}
