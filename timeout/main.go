package main

import (
	"fmt"
	"time"
)

func main() {
	c := boring("boring!")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
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
			time.Sleep(time.Duration(1 * time.Second))
		}
	}()
	return c
}