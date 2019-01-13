package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c)
	}
	fmt.Println("You are boring.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintln(msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}