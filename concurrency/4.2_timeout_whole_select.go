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
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Time is up buddy!")
			return
		}
	}

	fmt.Println("I quit! Stop talking ..")
}
