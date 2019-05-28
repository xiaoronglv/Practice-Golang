package main

import (
	"fmt"
)

type order struct {
	product_id    int
	user_id       int
	price         int
	product_count int
}

func (order order) amount() int {
	return order.price * order.product_count
}

func main() {
	o := order{product_id: 3, user_id: 1312354, price: 3000, product_count: 3}
	fmt.Println(o.amount()) // 如果我忘记加括号，返回的值是 0x1091570，真是奇怪，难道是这个 function 的内存地址吗？

}
