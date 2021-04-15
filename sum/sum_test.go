package main

import "testing"

var testData = GenerateData()

// generate billion slice data
func GenerateData() []int {
	data := make([]int, 1000000000)
	for key, _ := range data {
		data[key] = key % 128
	}

	return data
}

func BenchmarkSum(t *testing.B) {
	var testResult int
	for i := 0; i < t.N; i++ {
		Sum(testData, &testResult)
	}
}

func BenchmarkSum1(t *testing.B) {
	var testResult int
	for i := 0; i < t.N; i++ {
		Sum1(testData, &testResult)
	}
}

func BenchmarkSum2(t *testing.B) {
	var testResult int
	for i := 0; i < t.N; i++ {
		Sum2(testData, &testResult)
	}
}

func BenchmarkSum3(t *testing.B) {
	var testResult int
	for i := 0; i < t.N; i++ {
		Sum3(testData, &testResult)
	}
}
