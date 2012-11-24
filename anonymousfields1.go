package main
import "fmt"
type Kitchen struct {
    numOfPlates int
}
type House struct {
    Kitchen // 匿名字段
    numOfRooms int
}
func main() {
    h := House{Kitchen{10}, 3} // 初始化需要使用Struct名
    fmt.Println("房屋 h 有", h.numOfRooms, "间房") // numOfRooms 是 House 的一个字段
    fmt.Println("房屋 h 有", h.numOfPlates, "个盘子") // numOfPlates 是Kitchen提供的字段，
                                                      //而Kitchen又是House的一个匿名字段，所以这里可以访问到它
    fmt.Println("这间房屋的厨房有：", h.Kitchen) // 我们可以通过Struct的名称访问整个匿名Struct的内容
}
