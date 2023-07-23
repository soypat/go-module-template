package gomoduletemplate

// Fibonacci returns the nth number in the Fibonacci sequence.
func Fibonacci(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
