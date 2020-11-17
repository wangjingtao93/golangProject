package main
import (
	"bytes"
	"fmt"
)

//func Contains(b, subslice []byte) bool
//bytes.Constains 包含报告子切片是否在b内。
func main()  {
	var s1,s2 []byte
	s1 = []byte("abcfoo")
	s2 = []byte("abc")
	if bytes.Contains(s1,s2) {
	   fmt.Println(" s1 constains s2 ")
	}
  
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}