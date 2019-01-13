package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := fanIn(boring("boring c!"), boring("boring d!"))
	for i := 0; i < 15; i++ {
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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {for {c <- <-input1}}()
	go func() {for {c <- <-input2}}()
	return c
}
