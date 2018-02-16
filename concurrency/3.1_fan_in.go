package main

import (
	"fmt"
	"math/rand"
	"time"
)

func speak(str string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", str, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	poorva := speak("poorva")
	mohan := speak("mohan")

	for i := 0; i < 5; i++ {
		fmt.Printf("Speaking %q\n", <-poorva)
		fmt.Printf("Speaking %q\n", <-mohan)
	}

	fmt.Println("I quit! Stop talking ..")
}
