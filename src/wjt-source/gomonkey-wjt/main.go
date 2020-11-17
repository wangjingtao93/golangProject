package main

import (
	"fmt"
	"wjt-source/gomonkey-wjt/test/fake"
	"reflect"
)

func main() {
	output, err := fake.ExecWjt("", "")
	if err == nil {
		fmt.Println(output)
	}
	//strList := []string{"abc"}
	flag := fake.Belong("", nil)
	fmt.Println(flag)

    slice := fake.NewSlice()
	var s *fake.Slice
	m := reflect.TypeOf(s)
	fmt.Println(m.Elem)

	slice.Add(1)



}
