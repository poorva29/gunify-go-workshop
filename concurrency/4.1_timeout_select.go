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
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := speak("poorva")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Speak up fast!")
			return
		}
	}

	fmt.Println("I quit! Stop talking ..")
}
