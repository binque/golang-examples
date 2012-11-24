package main
import "fmt"
func main() {
    for i := 0; i < 11; i++ {
        if i%2 == 0 {
            continue //我不输入偶数
        }
        fmt.Println("当前i的值为：", i)
    }
}
