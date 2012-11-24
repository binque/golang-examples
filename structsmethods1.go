package main
import "fmt"
type Rectangle struct {
    width, length int
}
func (r Rectangle) Area() int {
    return r.width * r.length
}
func main() {
    r := Rectangle{4,3}
    fmt.Println("Rectangle is:", r)
    fmt.Println("Rectangle area is:", r.Area())
}

