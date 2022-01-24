/**
 * @Author: Seedzheng
 */

package templatemethod

import "testing"

func TestFruit_Eat(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "template method",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fruit{
				FruitEatPrepare: CutPeel{},
			}
			f.Eat()
		})
	}
}
