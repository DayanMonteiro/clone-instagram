package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Olá Mundo!") // goroutine
	write("Programando em Go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
