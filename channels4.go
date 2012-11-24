package main

import (
	"fmt"
	"time"
	"strconv"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName // 传递一个 cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs, "Strawberry", 5)
	go makeCakeAndSend(cs, "Chocolate", 10)
	go receiveCakeAndPack(cs)

	du,_ := time.ParseDuration("2s")
	time.Sleep(du)
}
