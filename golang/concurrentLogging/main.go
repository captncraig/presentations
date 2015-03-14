package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
//buffer can hold 100 messages before writer blocks
var logChan = make(chan string, 100)
var wg = sync.WaitGroup{}

func main() {
	go logToConsole()
	wg.Add(3)
	go doSomething(1)
	go doSomething(2)
	go doSomething(3)
	wg.Wait()
}
func logToConsole() {
	for {
		fmt.Println(<-logChan)
	}
}
func doSomething(i int) {
	defer wg.Done()
	//do real work
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	logChan <- fmt.Sprintf("I did something %d", i)
}

// END OMIT
