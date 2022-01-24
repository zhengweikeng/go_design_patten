/**
 * @Author: Seedzheng
 */

package factorygen

import "fmt"

type Human interface {
    Speak()
}

type HumanType uint
type HumanFactory func() Human

const (
    HumanTypeChinese HumanType = iota
    HumanTypeJapanese
    HumanTypeAmerican
)

type Chinese struct {
}

func (c Chinese) Speak() {
    fmt.Println("I speak Chinese")
}

func produceChinese() Human {
    return &Chinese{}
}

type Japanese struct {
}

func (j Japanese) Speak() {
    fmt.Println("I speak Japanese")
}

func produceJapanese() Human {
    return &Japanese{}
}

type American struct {
}

func (a American) Speak() {
    fmt.Println("I speak American English")
}

func produceAmerican() Human {
    return &American{}
}

var humanFactoryGenMap = map[HumanType]HumanFactory{
    HumanTypeChinese:  produceChinese,
    HumanTypeJapanese: produceJapanese,
    HumanTypeAmerican: produceAmerican,
}

func GetHumanFactory(humanType HumanType) HumanFactory {
    if h, ok := humanFactoryGenMap[humanType]; ok {
        return h
    }

    return nil
}

func HumanSpeak(humanType HumanType) {
    humanFactory := GetHumanFactory(humanType)
    if humanFactory == nil {
        return
    }

    human := humanFactory()
    human.Speak()
}
