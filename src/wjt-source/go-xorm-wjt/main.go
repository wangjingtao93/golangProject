package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "wjt-source/go-xorm-wjt/xorm_models/models"
    "time"
)

func main() {

    var engine *xorm.Engine
    //连接数据库
    engine, err := xorm.NewEngine("mysql", "root:root@tcp(192.168.56.130:3306)/test?charset=utf8")
    if err != nil {
        fmt.Println(err)
        return
    }
    //连接测试
    if err := engine.Ping(); err != nil {
        fmt.Println(err)
        return
    }
    defer engine.Close() //延迟关闭数据库
    fmt.Println("数据库链接成功")

    //查询单条数据
    var doc models.DoctorTb
    b, _ := engine.Where("name = ?", "钟南山").Get(&doc)
    if b {
        fmt.Println(doc)
    } else {
        fmt.Println("数据不存在")
    }

    //查询单条数据方式2 会根据结构体的
    doc2 := models.DoctorTb{Name: "钟南山"}
    b, _ = engine.Get(&doc2)
    fmt.Println(doc2)

    //新增数据
    doc3 := models.DoctorTb{0, "王医生", 48, 1, time.Now()}
    i3, _ := engine.InsertOne(doc3)
    fmt.Println("新增结果：", i3)

    //查询列表
    docList := make([]models.DoctorTb, 0)
    engine.Where("age > ? or name like ?", 40, "林%").Find(&docList)
    fmt.Println("docList：", docList)

    //查询列表方式2
    docList2 := make([]models.DoctorTb, 0)
    engine.Where("age > ?", 40).Or("name like ?", "林%").OrderBy("Id desc").Find(&docList2)
    fmt.Println("docList2：", docList2)

    //查询分页
    docList3 := make([]models.DoctorTb, 0)
    page := 0     //页索引
    pageSize := 2 //每页数据
    limit := pageSize
    start := page * pageSize
    totalCount, err := engine.Where("age > ? or name like ?", 40, "林%").Limit(limit, start).FindAndCount(&docList3)
    fmt.Println("总记录数：", totalCount, "docList3：", docList3)

    //直接用语句查询
    docList4 := make([]models.DoctorTb, 0)
    engine.SQL("select * from doctor_tb where age > ?", 40).Find(&docList4)
    fmt.Println("docList4：", docList4)

    //删除
    docDel := models.DoctorTb{Name: "王医生"}
    iDel, _ := engine.Delete(&docDel)
    fmt.Println("删除结果：", iDel)

    //删除方式2
    engine.Exec("delete from doctor_tb where Id = ?", 3)

    //更新数据
    doc5 := models.DoctorTb{Name: "钟医生"}
    //更新数据ID为2的记录名字更改为“钟医生”
    iUpdate, _ := engine.Id(2).Update(&doc5)
    fmt.Println("更新结果：", iUpdate)

    //指定表名查询.Table()
    user := models.UserTb{Id: 2}
    b, _ = engine.Table("user_tb").Get(&user)
    fmt.Println(user)

    //事务
    session := engine.NewSession()
    defer session.Close()
    err = session.Begin()
    _, err = session.Exec("delete from doctor_tb where Id = ?", 6)
    if err != nil {
        session.Rollback()
        return
    }
    _, err = session.Exec("delete from user_tb where Id = ?", 10)
    if err != nil {
        session.Rollback()
        return
    }
    err = session.Commit()
    if err != nil {
        return
    }
    fmt.Println("事务执行成功")
}