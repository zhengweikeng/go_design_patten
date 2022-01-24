/**
 * @Author: Seedzheng
 * 中介模式
 */

package mediator

import "fmt"

type SignalType int

const (
	AirPosition SignalType = iota
	IsAllowLanding
)

type AirTower struct {
}

func (tower AirTower) ReceiveSignal(signalType SignalType) {
	if signalType == AirPosition {
		fmt.Println("发送飞机方位")
	} else if signalType == IsAllowLanding {
		fmt.Println("飞机允许下降")
	}
}

type Airplane struct {
}

func (airplane Airplane) ReceiveAirPos() {
	tower := AirTower{}
	tower.ReceiveSignal(AirPosition)
}
