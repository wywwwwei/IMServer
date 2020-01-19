package main

import (
	"github.com/wywwwwei/IMServer/Service"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		Service.UserInit()
		setupHttp()
		wg.Done()
	}()
	go func() {
		setTCP()
		wg.Done()
	}()

	wg.Wait()
}