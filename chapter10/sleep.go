package main

import (
	"fmt"
	"time"
)

func sleep(seconds float64) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	sleep(5)
	fmt.Println("Slept for a while")
}
