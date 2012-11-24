package main

import (
	"fmt"
	"time"
	"strconv"
)

var i int

func makeCakeAndSend(cs chan string) {
	i = i + 1
	cakeName := "Strawberry Cake " + strconv.Itoa(i)
	fmt.Println("Making a cake and sending...", cakeName)
	cs <- cakeName // send a strawberry cake
}

func receiveCakeAndPack(cs chan string) {
	s := <- cs // get whatever cake is on the channel
	fmt.Println("Packing received cake: ", s)
}

func main() {
	cs := make(chan string)
	for i := 1; i < 5; i++ {
		go makeCakeAndSend(cs)
		go receiveCakeAndPack(cs)
	}
	// Sleep for a while so that the program doesn't exit immediately and output is clear for illustration
	wait, _ := time.ParseDuration("2s")
	time.Sleep(wait)
}
