package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Go routines change
	//Old way
	wg := sync.WaitGroup{}
	wg.Add(1)
	go doSomeGoRoutine(&wg)
	wg.Wait()

	//New way
	wg.Go(doSomeGoRoutineNew)
	wg.Wait()
}

func doSomeGoRoutine(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		wg.Done()
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Hello from goroutine")

}

func doSomeGoRoutineNew() {
	time.Sleep(time.Second * 2)
	fmt.Println("Hello from new goroutine")
}
