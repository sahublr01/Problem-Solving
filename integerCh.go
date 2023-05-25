package main

import "fmt"

func sendNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sendNumbers(ch)

	for num := range ch {
		fmt.Println(num)
	}
}
