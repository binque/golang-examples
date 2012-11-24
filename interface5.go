package main
import "fmt"
// Go第一步：定义你的数据结构
type Bus struct {
    l, b, h int
    rows, seatsPerRow int
}
// Go第二步：定义一个现实世界对象的抽象层
type Cuboider interface {
    CubicVolume() int
}
// Go第三步：实现接口的方法
func (bus Bus) CubicVolume() int {
    return bus.l * bus.b * bus.h
}
// 重复上面的二三步
type PublicTransporter interface {
    PassengerCapacity() int
}
func (bus Bus) PassengerCapacity() int {
    return bus.rows * bus.seatsPerRow
}
// 新的项目需求
type PersonalSpaceLaw interface {
    IsCampliantWithLaw() bool
}
func (bus Bus) IsCampliantWithLaw() bool {
    return ( bus.l * bus.b * bus.h ) / (bus.rows * bus.seatsPerRow) >= 3
}
func main() {
    bus := Bus{
        l:10, b:6, h:3,
        rows: 10, seatsPerRow: 5}
    fmt.Println("Cubic volume of bus: ", bus.CubicVolume())
    fmt.Println("Maximum number of passengers:", bus.PassengerCapacity())
    fmt.Println("Is campliant with law:", bus.IsCampliantWithLaw())
}
