package main
import "fmt"
type Kitchen struct {
    numOfLamps int;
}
type Bedroom struct {
    numOfLamps int;
}
type House struct {
    Kitchen
    Bedroom
}
func main() {
    h := House{Kitchen{2}, Bedroom{3}} // Kitchen 2, Bedroom 3
    fmt.Println("Ambiguous number of lamps:", h.numOfLamps) // 这会出错
}
