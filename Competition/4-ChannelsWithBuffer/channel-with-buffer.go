package main

import "fmt"

func main() {
	channel := make(chan string, 200)
	channel <- "Olá Mundo!"
	channel <- "Programando em Go!"
	channel <- "Programando em Go De Novo!"

	message := <-channel
	message2 := <-channel
	message3 := <-channel

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
}
