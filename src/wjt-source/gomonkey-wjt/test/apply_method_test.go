package test

import (
    "github.com/agiledragon/gomonkey"
    "github.com/smartystreets/goconvey/convey"
    "testing"
    "reflect"
    "wjt-source/gomonkey-wjt/test/fake"
)


func TestApplyMethod(t *testing.T) {
    slice := fake.NewSlice()
    var s *fake.Slice
    convey.Convey("TestApplyMethod", t, func() {

        convey.Convey("for succ", func() {
            err := slice.Add(1)
            convey.So(err, convey.ShouldEqual, nil)
            // 第一个参数是目标类的指针变量的反射类型，第二个参数是字符串形式的方法名，第三个参数是桩函数
            patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
                return nil
            })
            defer patches.Reset()
            err = slice.Add(1)
            convey.So(err, convey.ShouldEqual, nil)
            err = slice.Remove(1)
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(len(slice), convey.ShouldEqual, 0)
        })

        convey.Convey("for already exist", func() {
            err := slice.Add(2)
            convey.So(err, convey.ShouldEqual, nil)
            patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
                return fake.ERR_ELEM_EXIST
            })
            defer patches.Reset()
            err = slice.Add(1)
            convey.So(err, convey.ShouldEqual, fake.ERR_ELEM_EXIST)
            err = slice.Remove(2)
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(len(slice), convey.ShouldEqual, 0)
        })

        convey.Convey("two methods", func() {
            err := slice.Add(3)
            convey.So(err, convey.ShouldEqual, nil)
            defer slice.Remove(3)
            patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
                return fake.ERR_ELEM_EXIST
            })
            defer patches.Reset()
            patches.ApplyMethod(reflect.TypeOf(s), "Remove", func(_ *fake.Slice, _ int) error {
                return fake.ERR_ELEM_NT_EXIST
            })
            err = slice.Add(2)
            convey.So(err, convey.ShouldEqual, fake.ERR_ELEM_EXIST)
            err = slice.Remove(1)
            convey.So(err, convey.ShouldEqual, fake.ERR_ELEM_NT_EXIST)
            convey.So(len(slice), convey.ShouldEqual, 1)
            convey.So(slice[0], convey.ShouldEqual, 3)
        })

        convey.Convey("one func and one method", func() {
            err := slice.Add(4)
            convey.So(err, convey.ShouldEqual, nil)
            defer slice.Remove(4)
            patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
                return outputExpect, nil
            })
            defer patches.Reset()
            patches.ApplyMethod(reflect.TypeOf(s), "Remove", func(_ *fake.Slice, _ int) error {
                return fake.ERR_ELEM_NT_EXIST
            })
            output, err := fake.Exec("", "")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, outputExpect)
            err = slice.Remove(1)
            convey.So(err, convey.ShouldEqual, fake.ERR_ELEM_NT_EXIST)
            convey.So(len(slice), convey.ShouldEqual, 1)
            convey.So(slice[0], convey.ShouldEqual, 4)
        })
    })
}

