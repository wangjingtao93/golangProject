package test

import (
    "github.com/agiledragon/gomonkey"
    "github.com/smartystreets/goconvey/convey"
    "testing"
    "wjt-source/gomonkey-wjt/test/fake"
)

func TestApplyFuncSeq(t *testing.T) {
    convey.Convey("TestApplyFuncSeq", t, func() {

        convey.Convey("default times is 1", func() {
            info1 := "hello cpp"
            info2 := "hello golang"
            info3 := "hello gomonkey"
            outputs := []gomonkey.OutputCell{
                {Values: gomonkey.Params{info1, nil}},// 模拟函数的第1次输出
                {Values: gomonkey.Params{info2, nil}},// 模拟函数的第2次输出
                {Values: gomonkey.Params{info3, nil}},// 模拟函数的第3次输出
            }
            //第一个参数是函数名，第二个参数是特定的桩序列参数
            patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
            defer patches.Reset()
            output, err := fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, info1)
            output, err = fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, info2)
            output, err = fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, info3)
        })

        convey.Convey("retry succ util the third times", func() {
            info1 := "hello cpp"
            outputs := []gomonkey.OutputCell{
                {Values: gomonkey.Params{"", fake.ErrActual}, Times: 2},// 模拟函数的第1次输出
                {Values: gomonkey.Params{info1, nil}},// 模拟函数的第2次输出
            }
            patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
            defer patches.Reset()
            output, err := fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, fake.ErrActual)
            output, err = fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, fake.ErrActual)
            output, err = fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, info1)
        })

        convey.Convey("batch operations failed on the third time", func() {
            info1 := "hello gomonkey"
            outputs := []gomonkey.OutputCell{
                {Values: gomonkey.Params{info1, nil}, Times: 2},
                {Values: gomonkey.Params{"", fake.ErrActual}},
            }
            patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
            defer patches.Reset()
            output, err := fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, info1)
            output, err = fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, nil)
            convey.So(output, convey.ShouldEqual, info1)
            output, err = fake.ReadLeaf("")
            convey.So(err, convey.ShouldEqual, fake.ErrActual)
        })

    })
}

