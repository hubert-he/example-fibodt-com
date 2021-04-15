package main

import "testing"

func TestFib1(t *testing.T) {
	for caseId, testCase := range []struct {
		num  int
		want int
	}{
		{1, 1},
		{2, 1},
		{5, 5},
		{6, 8},
		{9, 34},
	} {
		result := Fib1(testCase.num)
		if result != testCase.want {
			t.Errorf("case-%d Failed: fibo(%d)=%d, want %d", caseId, testCase.num, result, testCase.want)
			return
		}
	}
}

func TestFib2(t *testing.T) {
	for caseId, testCase := range []struct {
		num  int
		want int
	}{
		{1, 1},
		{2, 1},
		{5, 5},
		{6, 8},
		{9, 34},
	} {
		result := Fib2(testCase.num)
		if result != testCase.want {
			t.Errorf("case-%d Failed: fibo(%d)=%d, want %d", caseId, testCase.num, result, testCase.want)
			return
		}
	}
}

func BenchmarkFib1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Fib1(9)
	}
}

func BenchmarkFib2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Fib2(9)
	}
}

func BenchmarkFib3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Fib3(9)
	}
}
