package main

import "fmt"

func main() {

	ch := make(chan any)

	go func() {
		val, ok := <-ch

		fmt.Println(val, ok)
	}()
	ch <- "hello"
	close(ch)

}
