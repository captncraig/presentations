package main

// START OMIT
func hello(c chan<- string) {
	c <- "Hello, "
}
func world(c chan<- string) {
	c <- "World!"
}
func main() {
	c := make(chan string)
	go hello(c)
	go world(c)
	print(<-c)
	println(<-c)
}

// END OMIT
