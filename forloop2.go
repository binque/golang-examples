package main
import "fmt"
func main() {
//    for i := 0; ; i++ {
//        fmt.Println("i现在的值为：", i)
//    }
//    for i := 0; i < 3; {
//        fmt.Println("i现在的值为", i)
//    }
    s := ""
    for ; s != "aaaaa"; {
        fmt.Println("s的值为：", s)
        s = s + "a"
    }
}
