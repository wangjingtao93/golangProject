package myPackage

import (
	"fmt"
)


type myint int

//乘2

func (p *myint) mydouble() int {
	*p = *p * 2
	return 0
}

//平方
func (p myint) mysquare() int {
	p = p * p
	fmt.Println("mysquare p = ", p)
	return 0
}