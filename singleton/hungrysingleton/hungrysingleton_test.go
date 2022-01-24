/**
 * @Author: Seedzheng
 */

package hungrysingleton

import "testing"

func Test_logger_Info(t *testing.T) {
    type args struct {
        a []interface{}
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "Info",
            args: args{a: []interface{}{"test hungry singleton"}},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l := GetLogger()
            l.Info(tt.args.a...)
        })
    }
}
