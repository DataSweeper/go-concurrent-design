package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	go boring("boring!")
	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring.")
}
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)))
	}
}