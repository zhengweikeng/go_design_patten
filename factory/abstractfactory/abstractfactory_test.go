/**
 * @Author: Seedzheng
 */

package abstractfactory

import "testing"

func TestChinese_Speak(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "abstract factory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ChineseFactory{}
			c := f.ProduceHuman()
			c.Speak()
		})
	}
}
