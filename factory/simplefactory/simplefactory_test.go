/**
 * @Author: Seedzheng
 */

package simplefactory

import (
    "reflect"
    "testing"
)

func TestHumanFactory_ProduceHuman(t *testing.T) {
    type args struct {
        humanType HumanType
    }
    tests := []struct {
        name string
        args args
        want Human
    }{
        {
            name: "Simple factory",
            args: args{HumanTypeChinese},
            want: &Chinese{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := HumanFactory{}
            if got := f.ProduceHuman(tt.args.humanType); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ProduceHuman() = %v, want %v", got, tt.want)
            }
        })
    }
}
