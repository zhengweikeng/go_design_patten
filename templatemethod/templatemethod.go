/**
 * @Author: Seedzheng
 * 模板方法模式
 */

package templatemethod

import "fmt"

type FruitEatPrepare interface {
    BeforeEat()
}

type Fruit struct {
    FruitEatPrepare
}

func (f Fruit) Eat() {
    f.BeforeEat()
    fmt.Println("吃水果")
}

type CutPeel struct {
}

func (c CutPeel) BeforeEat() {
    fmt.Println("先削皮")
}
