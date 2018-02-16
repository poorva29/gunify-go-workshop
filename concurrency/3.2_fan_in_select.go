package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func speak(str string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", str, i)
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := fanIn(speak("poorva"), speak("mohan"))

	for i := 0; i < 10; i++ {
		fmt.Printf("Speaking %q\n", <-c)
	}

	fmt.Println("I quit! Stop talking ..")
}
