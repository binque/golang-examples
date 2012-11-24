package main
import "fmt"
// Shaper是一个接口，它只有一个函数 Area返回一个int类型的结果
type Shaper interface {
    Area() int
}
type Rectangle struct {
    length, width int
}
// 该 Area 函数工作于 Rectangle 类型，同时与 Shper有同样的定义，现在就称为Rectangle实现了Shaper接口
func (r Rectangle) Area() int {
    return r.length * r.width
}
type Square struct {
    side int
}
func (sq Square) Area() int {
    return sq.side * sq.side
}
func main() {
    r := Rectangle{length: 5, width: 3}
    fmt.Println("Rectangle r details are: ", r)
    fmt.Println("Rectangle r's area is: ", r.Area())
    s := Shaper(r)
    fmt.Println("Area of the Shape r is: ", s.Area())
    sq := Square{5}
    fmt.Println("Square details is: ", sq)
    fmt.Println("Area of square sq is: ", sq.Area())
    sqs := Shaper(sq)
    fmt.Println("Shaper sqs area is: ", sqs.Area())
}
