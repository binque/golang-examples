package main
import "fmt"
type Camera struct {}
func (_ Camera) takePicture() string {
    return "Clicked"
}
type Phone struct {}
func (_ Phone) call() string {
    return "Ring Ring"
}
type CameraPhone struct {
    Camera
    Phone
}
func main() {
    cp := new(CameraPhone)
    fmt.Println("A CameraPhone can take a picture: ", cp.takePicture())
    fmt.Println("A CameraPhone also can make a call: ", cp.call())
}
