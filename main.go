package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chs := make([]chan struct{}, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	strArr := [3]string{"a", "b", "c"}
	for i, v := range strArr {
		vc := v
		ch := make(chan struct{})
		chs[i] = ch
		go runRoutine(ch, &wg, vc)
	}
	time.Sleep(2 * time.Millisecond)
	for _, c := range chs {
		close(c)
		time.Sleep(2 * time.Millisecond)
	}
	wg.Wait()

}

func runRoutine(ch <-chan struct{}, wg *sync.WaitGroup, s string) {

	for {
		fmt.Printf("doing: %v\n", s)

		if isClosed(ch) {
			fmt.Printf("done: %v\n", s)
			wg.Done()
			return
		}
	}
}
func isClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
