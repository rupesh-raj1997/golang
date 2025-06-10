package main

func concurrentFib(n int) []int {
	ch := make(chan int, n)
	fibs := make([]int, 0, n) // this ensures 0 items in the array

	go fibonacci(n, ch)
	for item := range ch {
		fibs = append(fibs, item)
	}
	return fibs
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
