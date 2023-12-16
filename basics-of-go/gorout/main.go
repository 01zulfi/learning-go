package main

import "time"

func printMessage(string string, channel chan string) {
	for i := 0; i < 2; i++ {
		println(string)
		time.Sleep(time.Second)
	}
	channel <- "done"
}

func main() { // main goroutine
	// var channel chan string
	channel := make(chan string)

	go printMessage("go", channel) // new goroutine
	// go printMessage("non")
	// go printMessage("lang")
	// time.Sleep(time.Second * 10)

	response := <-channel
	println(response)
	close(channel)
}
