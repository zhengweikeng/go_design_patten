/**
 * @Author: Seedzheng
 */

package state

import (
	"fmt"
	"testing"
)

func TestSaiyanStateChange_ChangeToSuperSaiyan3(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "state",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ssc := NewSaiyanStateChange(1000)
			fmt.Println("当前赛亚人阶段:", ssc.saiyan.GetState(), "最低战斗值:", ssc.saiyan.GetMinFightPoint())
			ssc.ChangeToSuperSaiyan3()
			fmt.Println("当前赛亚人阶段:", ssc.saiyan.GetState(), "最低战斗值:", ssc.saiyan.GetMinFightPoint())
			ssc.ChangeToSuperSaiyan1()
			fmt.Println("当前赛亚人阶段:", ssc.saiyan.GetState(), "最低战斗值:", ssc.saiyan.GetMinFightPoint())
			ssc.ChangeToNormalSaiyan()
			fmt.Println("当前赛亚人阶段:", ssc.saiyan.GetState(), "最低战斗值:", ssc.saiyan.GetMinFightPoint())
		})
	}
}
