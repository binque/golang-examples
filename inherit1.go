package main
import "fmt"
type Car struct {
    wheelCount int
}
// 定义 Car 的行为
func (car Car) numberOfWheels() int {
    return car.wheelCount
}
type Ferrari struct {
    Car // 匿名字段 Car
}
func main() {
    f := Ferrari{Car{4}}
    fmt.Printf("一辆法拉利有 %d 个轮胎\n", f.numberOfWheels())
}
