package main

import (
	"fmt"
	"time"
	"math/rand"
)
	
func main() {
	c := boring("boring!")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <- timeout:
			fmt.Println("You are so slow.")
			return
		}
	}
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