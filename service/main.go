package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := boring("boring c!")
	d := boring("boring d!")
	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c)
		fmt.Printf("you say: %q\n", <-d)
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