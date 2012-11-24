package main
import "fmt"
// 定义生物这个接口
type Organism interface {
    // 它有一个方法 iLove
    iLove() string
}
// 定义人这种生物
type Human struct {}
// 定义人对生物这个接口的实现
func (human Human) iLove() string {
    return "做人该做的事情，我是一只人，不是一个狗！"
}
// 定义男人
type Man struct { Human }
func (man Man) iLove() string {
    return "钓鱼，我是一只男人。"
}
//重复上面的定义
type Woman struct { Human }
func (woman Woman) iLove() string {
    return "购物，我是一只女人。"
}
// 定义狗
type Dog struct {}
func (dog Dog) iLove() string {
    return "汪！汪汪！汪汪汪！我是一个狗！不是一只人。"
}
func main() {
    h := new (Human)
    m := new (Man)
    w := new (Woman)
    d := new (Dog)
    organisms := [...]Organism{h, m, w, d}
    for i, _ := range (organisms) {
        // 直接调用接口的 iLove 方法
        fmt.Println("我喜欢", organisms[i].iLove())
    }
}

