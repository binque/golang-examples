package main
import (
    "fmt"
    "time"
)
func simulateEvent(name string, timeInSecs string) {
    fmt.Println("Started ", name, ": Should take", timeInSecs, "seconds")
    du, _ := time.ParseDuration(timeInSecs)
    time.Sleep(du)
    fmt.Println("Finished ", name)
}
func main() {
    go simulateEvent("First", "10s")
    go simulateEvent("Second", "6s")
    go simulateEvent("Third", "3s")
    wait, _ := time.ParseDuration("15s")
    time.Sleep(wait)
    fmt.Println("Exited")
}
