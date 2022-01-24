/**
 * @Author: Seedzheng
 * 状态模式
 */

package state

// SaiyanState 赛亚人阶段
type SuperSaiyanState int

const (
	NomalSaiyanState  SuperSaiyanState = iota // 赛亚人普通阶段
	SuperSaiyanState1                         // 超级赛亚人1
	SuperSaiyanState2                         // 超级赛亚人2
	SuperSaiyanState3                         // 超级赛亚人3
)

// Saiyan 赛亚人
type Saiyan interface {
	GetState() SuperSaiyanState // 获取当前赛亚人阶段
	GetMinFightPoint() int      // 获取最低战斗点数
}

type NormalSaiyan struct {
	FightPoint int
}

func (saiyan *NormalSaiyan) GetMinFightPoint() int {
	return saiyan.FightPoint
}

func (saiyan *NormalSaiyan) GetState() SuperSaiyanState {
	return NomalSaiyanState
}

type SuperSaiyan1 struct {
	FightPoint int
}

func (saiyan *SuperSaiyan1) GetMinFightPoint() int {
	return saiyan.FightPoint + 10000
}

func (saiyan *SuperSaiyan1) GetState() SuperSaiyanState {
	return SuperSaiyanState1
}

type SuperSaiyan2 struct {
	FightPoint int
}

func (saiyan *SuperSaiyan2) GetMinFightPoint() int {
	return saiyan.FightPoint + 200000
}

func (saiyan *SuperSaiyan2) GetState() SuperSaiyanState {
	return SuperSaiyanState2
}

type SuperSaiyan3 struct {
	FightPoint int
}

func (saiyan *SuperSaiyan3) GetMinFightPoint() int {
	return saiyan.FightPoint + 1000000
}

func (saiyan *SuperSaiyan3) GetState() SuperSaiyanState {
	return SuperSaiyanState3
}

type SuperSaiyanChange interface {
	GetSaiyan() Saiyan
	ChangeToNormalSaiyan() // 变为普通赛亚人
	ChangeToSuperSaiyan1() // 变为超级赛亚1
	ChangeToSuperSaiyan2() // 变为超级赛亚2
	ChangeToSuperSaiyan3() // 变为超级赛3
}

type SaiyanStateChange struct {
	minFightPoint int
	saiyan        Saiyan
}

func NewSaiyanStateChange(fightPoint int) *SaiyanStateChange {
	return &SaiyanStateChange{
		saiyan:        &NormalSaiyan{FightPoint: fightPoint},
		minFightPoint: fightPoint,
	}
}

func (ssc SaiyanStateChange) GetSaiyan() Saiyan {
	return ssc.saiyan
}

func (ssc *SaiyanStateChange) ChangeToNormalSaiyan() {
	if ssc.saiyan.GetState() == NomalSaiyanState {
		return
	} else {
		ssc.saiyan = &NormalSaiyan{FightPoint: ssc.minFightPoint}
	}
}

func (ssc *SaiyanStateChange) ChangeToSuperSaiyan1() {
	if ssc.saiyan.GetState() == NomalSaiyanState {
		ssc.saiyan = &SuperSaiyan1{FightPoint: ssc.saiyan.GetMinFightPoint()}
	} else if ssc.saiyan.GetState() == SuperSaiyanState1 {
		return
	} else {
		ssc.saiyan = &SuperSaiyan1{FightPoint: ssc.minFightPoint}
	}
}

func (ssc *SaiyanStateChange) ChangeToSuperSaiyan2() {
	if ssc.saiyan.GetState() == NomalSaiyanState {
		ssc.ChangeToSuperSaiyan1()
		ssc.saiyan = &SuperSaiyan1{FightPoint: ssc.saiyan.GetMinFightPoint()}
	} else if ssc.saiyan.GetState() == SuperSaiyanState1 {
		ssc.saiyan = &SuperSaiyan2{FightPoint: ssc.saiyan.GetMinFightPoint()}
	} else if ssc.saiyan.GetState() == SuperSaiyanState2 {
		return
	} else {
		ssc.saiyan = &SuperSaiyan1{FightPoint: ssc.minFightPoint}
	}
}

func (ssc *SaiyanStateChange) ChangeToSuperSaiyan3() {
	if ssc.saiyan.GetState() == NomalSaiyanState {
		ssc.ChangeToSuperSaiyan1()
		ssc.ChangeToSuperSaiyan2()
		ssc.saiyan = &SuperSaiyan3{FightPoint: ssc.saiyan.GetMinFightPoint()}
	} else if ssc.saiyan.GetState() == SuperSaiyanState1 {
		ssc.ChangeToSuperSaiyan2()
		ssc.saiyan = &SuperSaiyan3{FightPoint: ssc.saiyan.GetMinFightPoint()}
	} else if ssc.saiyan.GetState() == SuperSaiyanState2 {
		ssc.saiyan = &SuperSaiyan3{FightPoint: ssc.saiyan.GetMinFightPoint()}
	} else {
		return
	}
}
