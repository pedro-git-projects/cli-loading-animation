package main

import (
	"sync"

	"github.com/pedro-git-projects/cli-loading-animation/src/loader"
)

func main() {
	wg := &sync.WaitGroup{}
	quit := make(chan bool)
	wg.Add(1)
	go loader.DoSomething(quit)
	go loader.Spinner(quit, wg)
	wg.Wait()
}
