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
func databaseCall(c chan<- string) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	c <- "abc"
}
func cacheCall(c chan<- string) {
	time.Sleep(time.Duration(rand.Intn(40)) * time.Millisecond)
	c <- "def"
}
func main() {
	dbChan := make(chan string, 1)
	cacheChan := make(chan string, 1)
	go databaseCall(dbChan)
	go cacheCall(cacheChan)
	select {
	case d := <-dbChan:
		fmt.Println("DB wins!", d)
	case c := <-cacheChan:
		fmt.Println("Cache wins!", c)
	}
}

//END OMIT
