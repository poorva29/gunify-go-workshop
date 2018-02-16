package main

import (
	"fmt"
	"math/rand"
	"time"
)

func speak(str string) {
	for i := 0; ; i++ {
		fmt.Println(str, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	go speak("Hey")

	fmt.Println("I am listening")
	time.Sleep(3 * time.Second)
	fmt.Println("I quit! Stop talking ..")
}
