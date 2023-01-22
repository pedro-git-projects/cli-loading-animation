package loader

import (
	"fmt"
	"sync"
	"time"
)

func DoSomething(done chan<- bool) {
	time.Sleep(5 * time.Second)
	done <- true
}

func Spinner(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("\nDone!\n")
	for {
		select {
		case <-done:
			return
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
