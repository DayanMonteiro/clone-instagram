package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexar(write("Olá Mundo!"), write("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexar(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case mensagem := <-inputChannel2:
				outputChannel <- mensagem
			}
		}
	}()

	return outputChannel

}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
