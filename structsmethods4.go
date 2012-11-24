package main
import "fmt"
import "time"
func (t time.Time) first3Chars() string {
    return t.Weekday().String()[0:3]
}
func main() {
    t := time.Time{}
    fmt.Println("First 3 Chars is: ", t.first3Chars())
}
