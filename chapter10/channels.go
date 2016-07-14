package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("Sending...")
		c <- "ping"
		fmt.Println("Sent...")
	}
}

func printer(c chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(30)))
		fmt.Println(<-c)
	}
}

func main() {
	var c chan string = make(chan string, 1)
	go pinger(c)
	go pinger(c)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
