package main
import "C"
import "fmt"
// 导出函数（必须是大写且带上 //export）
//export Add
func Add(a, b int) int {
    return a + b
}
//export Hello
func Hello() {
    fmt.Println("Hello from Go!")
}
// go build -o mymodule.so -buildmode=c-shared main.go
func main() {} // 必须有 main，但不会被执行