/**
 * @Author: Seedzheng
 */

package facade

import "testing"

func TestShoppingCar_Buy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "facade",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := BookOrder{}
			inv := BookInventory{}
			n := BookOrderNotify{}

			c := &ShoppingCar{
				ProductOrder: &o,
				OrderNotify:  &n,
				Inventory:    &inv,
			}
			c.Buy()
		})
	}
}
