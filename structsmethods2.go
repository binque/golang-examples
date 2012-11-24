package main
import "fmt"
type Rectangle struct {
    length, width int
}
func (r Rectangle) AreaByValue() int {
    return r.length * r.width
}
func (r *Rectangle) AreaByReference() int {
    return r.length * r.width
}
func main() {
    r := Rectangle{4,3}
    fmt.Println("Rectangle is: ", r)
    fmt.Println("Rectangle area is: ", r.AreaByValue())
    fmt.Println("Rectangle area is: ", r.AreaByReference())
    fmt.Println("Rectangle area is: ", (&r).AreaByValue())
    fmt.Println("Rectangle area is: ", (&r).AreaByReference())
}
