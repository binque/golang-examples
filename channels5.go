package main

import (
	"fmt"
	"time"
	"strconv"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName
	}
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false
	for {
		if (strbry_closed && choco_closed) { return }
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <- strbry_cs:
			if (!strbry_ok) {
				strbry_closed = true
				fmt.Println("... Strawberry channel closed!")
			} else {
				fmt.Println("Received from Strawberry channel. Now packing ", cakeName)
			}
		case cakeName, choco_ok := <- choco_cs:
			if (!choco_ok) {
				choco_closed = true
				fmt.Println("... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel. Now packing ", cakeName)
			}
		}
	}
}

func main() {
	scs := make(chan string)
	ccs := make(chan string)
	go makeCakeAndSend(scs, "Strawberry", 3)
	go makeCakeAndSend(ccs, "Chocolate", 3)
	go receiveCakeAndPack(scs,ccs)
	du, _ := time.ParseDuration("3s")
	time.Sleep(du)
}
