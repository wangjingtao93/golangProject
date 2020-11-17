package main

import (
	"fmt"

	"github.com/smokezl/govalidators"
)

type Class struct {
    Cid       int64  `validate:"required||integer=10000,_"`
    Cname     string `validate:"required||string=1,5||unique"`
    BeginTime string `validate:"required||datetime=H:i"`
}

type Student struct {
    Uid          int64    `validate:"required||integer=10000,_"`
    Name         string   `validate:"required||string=1,5"`
    Age          int64    `validate:"required||integer=10,30"`
    Sex          string   `validate:"required||in=male,female"`
    Email        string   `validate:"email||user||vm"`
    PersonalPage string   `validate:"url"`
    Hobby        []string `validate:"array=_,2||unique||in=swimming,running,drawing"`
    CreateTime   string   `validate:"datetime"`
    Class        []Class  `validate:"array=1,3"`
}

func main() {
	validator := govalidators.New()

	student := &Student{
		Uid:          1234567,
		Name:         "张三",
		Age:          31,
		Sex:          "male1",
		Email:        "@qq.com",
		PersonalPage: "www.abcd.com",
		Hobby:        []string{"swimming", "singing"},
		CreateTime:   "2018-03-03 05:60:00",
		Class: []Class{
			Class{
				Cid:       12345678,
				Cname:     "语文",
				BeginTime: "13:00",
			},
			Class{
				Cid:       22345678,
				Cname:     "数学",
				BeginTime: "13:00",
			},
			Class{
				Cid:       32345678,
				Cname:     "数学",
				BeginTime: "13:60",
			},
		},
	}
	errList := validator.Validate(student)
	if errList != nil {
		for _, err := range errList {
			fmt.Println("err:", err)
		}
	}
}
