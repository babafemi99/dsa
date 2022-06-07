package main

import "fmt"

func fib(n int) int {
	var t0 = 0
	var t1 = 1
	var s int
	var i int
	if n <= 1 {
		return n
	}
	for i = 2; i <= n; i++ {
		s = t0 + t1
		t0 = t1
		t1 = s
	}
	return s
}

var f = make([]int, 10)

func fibRecursion(n int) int {
	if n <= 1 {
		return n
	}
	if f[n-2] == 0 {
		f[n-2] = fibRecursion(n - 2)
	}
	if f[n-1] == 0 {
		f[n-1] = fibRecursion(n - 1)
	}
	return f[n-2] + f[n-1]
}
func main() {
	fmt.Println(fibRecursion(10))
}
