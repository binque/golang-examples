package main
import "fmt"
// 定义 Rectangle 结构
type Rectangle struct {
    length, width int
}
func main() {
    r := Rectangle{}
    fmt.Println("默认的长方形是：", r) // 打印出默认的 Rectangle 的归零值
}
