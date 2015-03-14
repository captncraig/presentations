package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func main() {
	dbChan := make(chan string, 1)
	cacheChan := make(chan string, 4)
	go databaseCall(dbChan)
	go cacheCall(cacheChan) // HL
	go cacheCall(cacheChan) // HL
	go cacheCall(cacheChan) // HL
	go cacheCall(cacheChan) // HL
	select {
	case d := <-dbChan:
		fmt.Println("DB wins!", d)
	case c := <-cacheChan:
		fmt.Println("Cache wins!", c)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("Everybody loses: Timeout")
	}
}

//END OMIT

func databaseCall(c chan string) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	c <- "abc"
}

func cacheCall(c chan string) {
	time.Sleep(time.Duration(rand.Intn(40)) * time.Millisecond)
	c <- "def"
}
