
package main
import (
    "encoding/json"
    "fmt"
)
 
// Product 商品信息
type Product struct {
    Name      string  `json:"name"`
    ProductID int64   `json:"-"` // 表示不进行序列化
    Number    int     `json:"number"`
    Price     float64 `json:"price"`
    IsOnSale  bool    `json:"is_on_sale,string"`
}
 
func main() {
    
	//序列化
    p := &Product{}
    p.Name = "Xiao mi 6"
    p.IsOnSale = true
    p.Number = 10000
    p.Price = 2499.00
    p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Println()

	//反序列化
	up := &Product{}
	err := json.Unmarshal([]byte(data), up)
    fmt.Println(err)
    fmt.Println(*up)

}