package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Olá Mundo!", channel)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensage := range channel {
		fmt.Println(mensage)
	}

	fmt.Println("Fim do programa!")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
