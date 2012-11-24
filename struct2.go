package main
import "fmt"
type Rectangle struct {
    length, width int
    name string
}
func main() {
    r1 := Rectangle{2, 1, "First Rectangle"}
    fmt.Println("第一个长方形 r1 为：", r1)
    r2 := Rectangle{width: 3, name: "Second Rectangle", length: 4}
    fmt.Println("第二个长方形 r2 为：", r2)
    pr := new (Rectangle) // 获取一个指向一个新的长方形的指针
    (*pr).width = 6 // 设置 pr 的宽为 6
    pr.length = 8 // 与上面是同样的效果，看这里是没有 -> 符号的
    pr.name = "指针长方形"
    fmt.Println("指针长方形是：", pr)
}
