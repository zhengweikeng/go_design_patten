/**
 * @Author: Seedzheng
 * 桥接模式
 * 例子：冲咖啡，可以选择大中小杯，也可选择去冰、加糖、加奶
 */

package bridge

import "fmt"

type CoffeeCupType uint8
type CoffeeAdditionType uint8

const (
    CoffeeCupTypeSmall  CoffeeCupType = iota // 小杯
    CoffeeCupTypeMedium                      // 中杯
    CoffeeCupTypeLarge                       // 大杯
)

const (
    CoffeeAdditionTypeNoIce CoffeeAdditionType = iota // 去冰
    CoffeeAdditionTypeSugar                           // 加糖
    CoffeeAdditionTypeMilk                            // 加奶
)

type Coffee interface {
    MakeCoffee()
}

type CoffeeAddition interface {
    AddAddition()
}

type SmallCoffee struct {
    CoffeeAddition
}

func (c SmallCoffee) MakeCoffee() {
    c.AddAddition()
    fmt.Println("冲小杯咖啡")
}

type MediumCoffee struct {
    CoffeeAddition
}

func (c MediumCoffee) MakeCoffee() {
    c.AddAddition()
    fmt.Println("冲中杯咖啡")
}

type LargeCoffee struct {
    CoffeeAddition
}

type NoIce struct {
}

func (a NoIce) AddAddition() {
    fmt.Println("咖啡去冰")
}

type Milk struct {
}

func (m Milk) AddAddition() {
    fmt.Println("咖啡加奶")
}

type Sugar struct {
}

func (m Sugar) AddAddition() {
    fmt.Println("咖啡加糖")
}

func (c LargeCoffee) MakeCoffee() {
    c.AddAddition()
    fmt.Println("冲大杯咖啡")
}

var coffeeAdditionFuncMap = map[CoffeeAdditionType]func() CoffeeAddition{
    CoffeeAdditionTypeNoIce: NewNoIce,
    CoffeeAdditionTypeMilk:  NewMilk,
    CoffeeAdditionTypeSugar: NewSugar,
}

func NewCoffeeAddition(additionType CoffeeAdditionType) CoffeeAddition {
    h, ok := coffeeAdditionFuncMap[additionType]
    if ok {
        return h()
    }

    return nil
}

func NewNoIce() CoffeeAddition {
    return &NoIce{}
}

func NewMilk() CoffeeAddition {
    return &Milk{}
}

func NewSugar() CoffeeAddition {
    return &Sugar{}
}

func NewSmallCoffee(additionType CoffeeAdditionType) Coffee {
    return &SmallCoffee{NewCoffeeAddition(additionType)}
}

func NewMediumCoffee(additionType CoffeeAdditionType) Coffee {
    return &MediumCoffee{NewCoffeeAddition(additionType)}
}

func NewLargeCoffee(additionType CoffeeAdditionType) Coffee {
    return &LargeCoffee{NewCoffeeAddition(additionType)}
}

var coffeeFuncMap = map[CoffeeCupType]func(CoffeeAdditionType) Coffee{
    CoffeeCupTypeSmall:  NewSmallCoffee,
    CoffeeCupTypeMedium: NewMediumCoffee,
    CoffeeCupTypeLarge:  NewLargeCoffee,
}

func NewCoffee(cupType CoffeeCupType, additionType CoffeeAdditionType) Coffee {
    h, ok := coffeeFuncMap[cupType]
    if ok {
        return h(additionType)
    }

    return nil
}
