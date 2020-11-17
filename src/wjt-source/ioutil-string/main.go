package main
 
import (
    "fmt"
    "io/ioutil"
    "strings"
)
 
func main() {
    r1 := strings.NewReader("aaa")
    //返回ReadCloser对象提供close函数
    rc1 := ioutil.NopCloser(r1)
    defer rc1.Close()
 
    //ReadAll读取所有数据
    p, _ := ioutil.ReadAll(strings.NewReader("12345"))
    fmt.Println(string(p))
 
    //ReadDir返回目录下所有文件切片
    fileInfo, _ := ioutil.ReadDir("/root")
    for _, data := range fileInfo {
        fmt.Println(data.Name())
    }
 
    //读取整个文件数据
    data, _ := ioutil.ReadFile("./1.rtf")
    fmt.Println(string(data))
 
    //创建文件，存在清空文件
    ioutil.WriteFile("./1.txt", []byte("111"), 0655)
 
    //创建指定前缀的临时文件夹,返回文件夹名称
    dir, _ := ioutil.TempDir("./", "test")
    fmt.Println(dir)
 
    //创建test为前缀的临时文件，返回os.File指针
    f, _ := ioutil.TempFile("./", "test")
    f.Write([]byte("222"))
    f.Close()
}