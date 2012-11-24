package main
import "fmt"
func main() {
    for i, j, s := 0,5,"a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s + "a" {
        fmt.Printf("i=%d, j=%d, s=%s\n", i, j, s)
    }
}
