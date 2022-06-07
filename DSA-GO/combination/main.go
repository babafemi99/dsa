package main

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return factorial(n-1) * n
}

func nCr(n, r int) int {
	return factorial(n) / (factorial(n-r) * factorial(r))
}

func nCrRecursion(n, r int) int {
	if r == 0 || n == r {
		return 1
	}
	return nCrRecursion(n-1, r-1) + nCrRecursion(n-1, r)
}

var f []int

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	println(nCrRecursion(5, 5))
}
