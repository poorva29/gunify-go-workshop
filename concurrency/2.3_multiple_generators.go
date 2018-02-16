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
	c := speak("Hey")

	for i := 0; i < 5; i++ {
		fmt.Printf("Speaking %q\n", <-c)
	}

	fmt.Println("I quit! Stop talking ..")
}
