package service

import (
	"testing"
	"fmt"
)

func TestEcec(t *testing.T) {

	output, err := Exec("ls", "-l")
	if err == nil {
		t.Log("Login test1 success")
		fmt.Println(output)
	} else {
		t.Error(err)
	}
}
