package main

import "fmt"
import "time"

func makeOddGenerator() func() uint {
	i := uint(1)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	odd := makeOddGenerator()
	for i := uint(0); i < 999; i = odd() {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}
