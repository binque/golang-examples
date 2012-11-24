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
    q := Square{side:5}
    shapes := [...]Shaper{r, q}
    fmt.Println("Looping through shapes for area ...")
    for n, _ := range shapes {
        fmt.Printf("Shape detail: %v , and is area is: %d\n", shapes[n], shapes[n].Area()) 
    }
}
