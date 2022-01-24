/**
 * @Author: Seedzheng
 * 门面模式（外观模式）
 */

package facade

import "fmt"

type ProductOrder interface {
	PlaceOrder()
}

type OrderNotify interface {
	SendUserMsg()
}

type Inventory interface {
	ChangeInventory()
}

type BookOrder struct {
}

func (o *BookOrder) PlaceOrder() {
	fmt.Println("书籍下单")
}

type BookOrderNotify struct {
}

func (o *BookOrderNotify) SendUserMsg() {
	fmt.Println("书籍下单通知")
}

type BookInventory struct {
}

func (o *BookInventory) ChangeInventory() {
	fmt.Println("书籍库存扣减")
}

type ShoppingCar struct {
	ProductOrder
	OrderNotify
	Inventory
}

func (c *ShoppingCar) Buy() {
	c.PlaceOrder()
	c.ChangeInventory()
	c.SendUserMsg()
}
