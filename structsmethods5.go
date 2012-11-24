package main
import (
    "fmt"
    "time"
)
type myTime struct {
    time.Time // 匿名字段
}
func (t myTime) first3Chars() string {
    return t.Weekday().String()[0:3]
}
func main() {
    m := myTime{}
    fmt.Println("Full Weekday:", m.Weekday().String())
    fmt.Println("First 3 chars:", m.first3Chars())
}
