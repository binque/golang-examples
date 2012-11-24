package main
import "fmt"
func LoopIt(times int) {
    for i := 0; i < times; i++ {
        fmt.Printf("%d\t", i)
    }
    fmt.Println("--Loop End--")
}
func main() {
    go LoopIt(10)
    LoopIt(10)
}
