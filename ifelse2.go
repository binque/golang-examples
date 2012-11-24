package main
import "fmt"
func main() {
    a, b := 4, 5
    if a < b {
        fmt.Println("a < b")
    } else if a > b {
        fmt.Println("a > b")
    } else {
        fmt.Println("a = b")
    }
}
