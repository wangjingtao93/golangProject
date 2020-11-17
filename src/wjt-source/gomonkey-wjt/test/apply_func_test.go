package test

import (
    "github.com/agiledragon/gomonkey"
    "github.com/smartystreets/goconvey/convey"
    "testing"
//    "encoding/json"
    "wjt-source/gomonkey-wjt/test/fake"
)

var (
    outputExpect = "xxx-vethName100-yyy"
)

func TestApplyFunc(t *testing.T) {
    convey.Convey("TestApplyFunc", t, func() {

        //wjt add
        convey.Convey("input and output param", func() {
            // 第一个参数是函数名，第二个参数是桩函数
            patches := gomonkey.ApplyFunc(fake.AddOne, func(t1 int32) int32 {
                return 0
            })//对函数AddOne打桩
            
            //patches 对象通过 Reset 成员方法删除所有测试桩。
            defer patches.Reset()
            patches.ApplyFunc(fake.MinusOne, func(t1 int32) int32 {
                return -1
            })//对函数MinusOne打桩
            result := fake.MultiAddOne(10) //看好了我调用的是MultiAddOne函数，而MultiAddOne函数内部调用了AddOne和MinusOne。
            convey.So(result, convey.ShouldEqual, 0)
            })

        convey.Convey("one func for succ", func() {
            patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
                    return outputExpect, nil
                })
            defer patches.Reset()
            output, err := fake.ExecWjt("","")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, outputExpect)
        })

        // convey.Convey("one func for fail", func() {
        //     patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
        //         return "", fake.ErrActual
        //     })
        //     defer patches.Reset()
        //     output, err := fake.Exec("", "")
        //     convey.So(err, convey.ShouldEqual, fake.ErrActual)
        //     convey.So(output, convey.ShouldEqual, "")
        // })

        // convey.Convey("two funcs", func() {
        //     patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
        //         return outputExpect, nil
        //     })
        //     defer patches.Reset()
        //     patches.ApplyFunc(fake.Belong, func(_ string, _ []string) bool {
        //         return true
        //     })
        //     output, err := fake.Exec("", "")
        //     convey.So(err, convey.ShouldEqual, nil)
        //     convey.So(output, convey.ShouldEqual, outputExpect)
        //     flag := fake.Belong("", nil)
        //     convey.So(flag, convey.ShouldBeTrue)
        // })

        // convey.Convey("input and output param", func() {
        //     patches := gomonkey.ApplyFunc(json.Unmarshal, func(_ []byte, v interface{}) error {
        //         p := v.(*map[int]int)
        //         *p = make(map[int]int)
        //         (*p)[1] = 2
        //         (*p)[2] = 4
        //         return nil
        //     })
        //     defer patches.Reset()
        //     var m map[int]int
        //     err := json.Unmarshal(nil, &m)
        //     convey.So(err, convey.ShouldEqual, nil)
        //     convey.So(m[1], convey.ShouldEqual, 2)
        //     convey.So(m[2], convey.ShouldEqual, 4)
        // })
    })
}
