package test

import (
    "github.com/agiledragon/gomonkey"
    "github.com/smartystreets/goconvey/convey"
    "testing"
    "github.com/agiledragon/gomonkey/test/fake"
)


func TestApplyFuncVar(t *testing.T) {
    convey.Convey("TestApplyFuncVar", t, func() {

        convey.Convey("for succ", func() {
            str := "hello"
            //第一个参数是函数变量的地址，第二个参数是桩函数
            patches := gomonkey.ApplyFuncVar(&fake.Marshal, func (_ interface{}) ([]byte, error) {
                return []byte(str), nil
            })
            defer patches.Reset()
            bytes, err := fake.Marshal(nil)
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(string(bytes), convey.ShouldEqual, str)
        })

        convey.Convey("for fail", func() {
            patches := gomonkey.ApplyFuncVar(&fake.Marshal, func (_ interface{}) ([]byte, error) {
                return nil, fake.ErrActual
            })
            defer patches.Reset()
            _, err := fake.Marshal(nil)
            convey.So(err, convey.ShouldEqual, fake.ErrActual)
        })
    })
}
