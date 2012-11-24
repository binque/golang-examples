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
// 定义只有 Ferrari 专属的方法
func (f Ferrari) sayHiToSchumacher() {
    fmt.Println("Hi Schumacher!")
}
type AstonMartin struct {
    Car
}
//添加只有 AstonMartin 专属的方法
func (a AstonMartin) sayHiToBond() {
    fmt.Println("Hi Bond, James Bond!")
}
func main() {
    f := Ferrari{Car{4}}
    fmt.Printf("一辆法拉利有 %d 个轮胎\n", f.numberOfWheels())
    f.sayHiToBond()
    a := AstonMartin{Car{4}}
    fmt.Printf("一辆阿斯顿马丁有 %d 个轮胎\n", a.numberOfWheels())
    a.sayHiToSchumacher()
}
