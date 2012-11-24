package main

import (
	"fmt"
	"time"
	"strconv"
)

func makeCakeAndSend(cs chan string) {
	for i := 1; i<=3; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		fmt.Println("Making a cake and sending...", cakeName)
		cs <- cakeName
	}
}

func receiveCakeAndPack(cs chan string) {
	for i := 1; i<=3; i++ {
		s := <- cs
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs)
	go receiveCakeAndPack(cs)

	wait, _ := time.ParseDuration("5s")
	time.Sleep(wait)
}
