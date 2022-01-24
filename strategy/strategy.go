/**
 * @Author: Seedzheng
 * 策略模式
 */

package strategy

import "fmt"

type CookStrategy interface {
	Cook()
}

type ChineseFood struct {
}

func (f ChineseFood) Cook() {
	fmt.Println("做中餐")
}

type FrenchFood struct {
}

func (f FrenchFood) Cook() {
	fmt.Println("做法餐")
}

type JapaneseFood struct {
}

func (f JapaneseFood) Cook() {
	fmt.Println("做日料")
}

type FoodType int

const (
	FoodTypeChinese FoodType = iota
	FoodTypeFrench
	FoodTypeJapanese
)

var foodMap = map[FoodType]CookStrategy{
	FoodTypeChinese:  ChineseFood{},
	FoodTypeFrench:   FrenchFood{},
	FoodTypeJapanese: JapaneseFood{},
}

func GetCookStrategy(foodType FoodType) CookStrategy {
	strategy, ok := foodMap[foodType]
	if ok {
		return strategy
	}

	return nil
}

type Dinner struct {
	DinnerType FoodType
}

func (d Dinner) CookDinner() {
	strategy := GetCookStrategy(d.DinnerType)
	if strategy == nil {
		return
	}

	strategy.Cook()
}
