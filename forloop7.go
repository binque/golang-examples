package main
import "fmt"
func main() {
    // 对于数组，range 返回数据的索引
    a := [...]string{"a", "b", "c", "d"}
    for i,v := range a {
        fmt.Println("数组的第", i+1, "个值为：", v)
    }

    // 对于 map，返回每一个键
    capitals := map[string]string {"France":"Paris", "Italy":"Rome", "China":"Beijing" }
    for key,value := range capitals {
        fmt.Println("Capital of", key, "is", value)
    }
}
