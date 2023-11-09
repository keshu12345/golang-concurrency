package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	userId := 100
	ch := make(chan any)

	var wg sync.WaitGroup

	go fetUserData(userId, ch, &wg)
	wg.Add(1)
	go findUserLikes(userId, ch, &wg)
	wg.Add(1)
	go UserRecommendation(userId, ch, &wg)
	wg.Add(1)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		select {
		case val, ok := <-ch:
			if !ok {
				break
			}
			fmt.Println(val)

		}

	}

	// for val := range ch {
	// 	fmt.Println(val)
	// }
}

func fetUserData(userId int, ch chan any, wg *sync.WaitGroup) {

	time.Sleep(150 * time.Millisecond)

	ch <- "user Data"
	wg.Done()

}

func findUserLikes(userId int, ch chan any, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)

	ch <- 110
	wg.Done()

}

func UserRecommendation(userId int, ch chan any, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)

	ch <- "user recommendation"
	wg.Done()
}
