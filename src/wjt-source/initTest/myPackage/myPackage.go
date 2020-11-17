package myPackage

import (
	"fmt"
)

type MP interface {
	mydouble() int
    mysquare() int 
}

var I int

func init() {
	I = 0
	fmt.Println("Call mypackage init1")
}

func init() {
	I = 1
	fmt.Println("Call mypackage init2")
}
