package main

import (
	"wjt-source/initTest/myPackage"
	"fmt"
)






func main() {
	fmt.Println("Hello go.... I = ", myPackage.I)
	var i myPackage.MP
	i.mydouble()
	fmt.Println("i = ", i)
	i.mysquare()
	fmt.Println("i = ", i)
}