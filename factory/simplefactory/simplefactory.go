/**
 * @Author: Seedzheng
 */

package simplefactory

import "fmt"

type Human interface {
    Speak()
}

type HumanType uint

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

type Japanese struct {
}

func (j Japanese) Speak() {
    fmt.Println("I speak Japanese")
}

type American struct {
}

func (a American) Speak() {
    fmt.Println("I speak American English")
}

type HumanFactory struct {
}

func (f HumanFactory) ProduceHuman(humanType HumanType) Human {
    if humanType == HumanTypeChinese {
        return &Chinese{}
    } else if humanType == HumanTypeJapanese {
        return &Japanese{}
    } else if humanType == HumanTypeAmerican {
        return &American{}
    } else {
        return nil
    }
}
