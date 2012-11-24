package main
import "fmt"
func main() {
    // 对于数组，range 返回数据的索引
    a := [...]string{"a", "b", "c", "d"}
    for i := range a {
        fmt.Println("数组的第", i+1, "个值为：", a[i])
    }

    // 对于 map，返回每一个键
    capitals := map[string]string {"France":"Paris", "Italy":"Rome", "China":"Beijing" }
    for key := range capitals {
        fmt.Println("Capital of", key, "is", capitals[key])
    }
}
