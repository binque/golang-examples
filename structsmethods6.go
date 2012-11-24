package main
import "fmt"
type Kitchen struct {
    numOfForks int
    numOfKnives int
}
func (k Kitchen) totalForksAndKnives() int {
    return k.numOfForks + k.numOfKnives
}
type House struct {
    Kitchen
}
func main() {
    h := House{Kitchen{4,4}}
    fmt.Println("Sum of forks and knives in house:", h.totalForksAndKnives())
}
