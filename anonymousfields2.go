package main
import "fmt"
type Kitchen struct {
    numOfLamps int
}
type House struct {
    Kitchen
    numOfLamps int
}
func main() {
    h := House{Kitchen{2}, 10} // Kitchen 有2个Lamps，House有10个
    fmt.Println("House有", h.numOfLamps, "个Lamps")
    fmt.Println("Kitchen有", h.Kitchen.numOfLamps, "个Lamps")
}
