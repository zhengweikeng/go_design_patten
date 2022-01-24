/**
 * @Author: Seedzheng
 */

package command

import "fmt"

type FightSkill interface {
    Spark()
}

type Chongjibo struct {
}

func (Chongjibo) Spark() {
    fmt.Println("发招:冲击波")
}

type Guipaiqigong struct {
}

func (Guipaiqigong) Spark() {
    fmt.Println("发招:龟派气功")
}

type Hero struct {
}

type FightRequest struct {
    SkillType int
}

func (h Hero) Fight(request FightRequest) {
    var skill FightSkill
    if request.SkillType == 0 {
        skill = Chongjibo{}
    } else if request.SkillType == 1 {
        skill = Guipaiqigong{}
    }

    skill.Spark()
}
