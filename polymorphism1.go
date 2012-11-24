package main
import "fmt"
// 定义人这个接口
type Human interface {
    // 它有一个方法 iLove
    iLove() string
}
// 定义男人
type Man struct {}
// 实现人这个接口
func (man Man) iLove() string {
    return "钓鱼"
}
//重复上面的定义
type Woman struct {}
func (woman Woman) iLove() string {
    return "购物"
}
func main() {
    m := new (Man)
    w := new (Woman)
    // 一个保存人的数组，它有两人个元素
    // 一个为 Man 类型，一个为Woman类型
    humans := [...]Human{m,w}
    for i, _ := range (humans) {
        // 直接调用接口的 iLove 方法
        fmt.Println("我喜欢", humans[i].iLove())
    }
}

