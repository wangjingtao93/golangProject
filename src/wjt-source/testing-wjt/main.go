package main

import (
	"fmt"
	"wjt-source/testing-wjt/service"
	"reflect"
)

func main() {
	output, err := service.ExecWjt("", "")
	if err == nil {
		fmt.Println(output)
	}
	//strList := []string{"abc"}
	flag := service.Belong("", nil)
	fmt.Println(flag)

    slice := service.NewSlice()
	var s *service.Slice
	m := reflect.TypeOf(s)
	fmt.Println(m.Elem)

	slice.Add(1)



}
