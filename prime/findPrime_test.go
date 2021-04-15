package main

import "testing"

func BenchmarkFindPrime(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FindPrime()
	}
}

func BenchmarkFindPrime1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FindPrime1()
	}
}

func BenchmarkFindPrime2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FindPrime2()
	}
}

func BenchmarkFindPrime3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FindPrime3()
	}
}

func BenchmarkFindPrime5(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FindPrime5()
	}
}
