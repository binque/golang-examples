package main
import "fmt"
func main() {
    i := 0
    for {
        if i >= 3 { break }
        fmt.Println("当前i的值为：", i)
        i++
    }
    fmt.Println("for后的第一个声明")
}
