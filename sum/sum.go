package main

func getLen(list []int) int {
	return len(list)
}

func Sum(data []int, result *int) {
	for i := 0; i < len(data); i++ {
		*result += data[i]
	}
}

func Sum1(data []int, result *int) {
	for i := 0; i < getLen(data); i++ {
		*result += data[i]
	}
}

func Sum2(data []int, result *int) {
	length := getLen(data)
	for i := 0; i < length; i++ {
		*result += data[i]
	}
}

func Sum3(data []int, result *int) {
	length := len(data)
	tmp := *result
	for i := 0; i < length; i++ {
		tmp += data[i]
	}
	*result = tmp
}
