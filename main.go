package main

import "fmt"

func main() {
	/*
		Unbuffered channel
	*/

	//ch := make(chan int)
	//fmt.Println("Main goroutine is running")
	//ch <- 1
	//
	//go func() {
	//	fmt.Println("Anonymous goroutine is running")
	//	fmt.Println(<-ch)
	//}()

	/*
		- Result: deadlock
		- Explain:
				1. ch := make(chan int): This creates an unbuffered channel ch that can send and receive integer values.
				2. fmt.Println("Main goroutine is running"): This prints a message indicating that the main goroutine has started executing.
				3. ch <- 1: This line attempts to send the value 1 into the channel ch.
			       However, since this is an unbuffered channel and there's no receiver ready yet.
			       So, this operation will block the main goroutine => the next commands are no execute

		- Fix:
		        1. Create another goroutine that precedes the operation ch <- 1
	*/

	//myCh := make(chan int)
	//fmt.Println("Main goroutine is running")
	//
	//go func() {
	//	fmt.Println("Anonymous goroutine is running")
	//	fmt.Println(<-myCh)
	//}()
	//
	//myCh <- 1

	ch := make(chan int)

	go func() {
		fmt.Println("Anonymous goroutine is running")
		fmt.Println("Get values from ch channel")
		for i := 0; i < 3; i++ {
			fmt.Println(<-ch)
		}
	}()

	fmt.Println("Main goroutine is running")
	ch <- 1
	ch <- 2
	ch <- 3
}
