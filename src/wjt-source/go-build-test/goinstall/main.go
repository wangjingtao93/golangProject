package main
import (
    "wjt-source/go-build-test/goinstall/mypkg"
    "fmt"
)
func main() {
    mypkg.CustomPkgFunc()
    fmt.Println("hello world")
}