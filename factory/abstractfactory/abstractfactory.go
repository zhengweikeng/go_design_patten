/**
 * @Author: Seedzheng
 */

package abstractfactory

import "fmt"

type Human interface {
	Speak()
}

type HumanCreateFactory interface {
	ProduceHuman()
}

type ChineseFactory struct {
}

func (f ChineseFactory) ProduceHuman() Human {
	return &Chinese{}
}

type Chinese struct {
}

func (c Chinese) Speak() {
	fmt.Println("I speak Chinese")
}

type JapaneseFactory struct {
}

func (f JapaneseFactory) ProduceHuman() Human {
	return &Japanese{}
}

type Japanese struct {
}

func (j Japanese) Speak() {
	fmt.Println("I speak Japanese")
}

type AmericanFactory struct {
}

func (f AmericanFactory) ProduceHuman() Human {
	return &American{}
}

type American struct {
}

func (a American) Speak() {
	fmt.Println("I speak American English")
}
