/**
 * @Author: Seedzheng
 * 单例饿汉模式
 */

package hungrysingleton

import "fmt"

type logger struct {
    Name string
}

func newLogger(name string) *logger {
    // ... 初始化logger
    return &logger{name}
}

var log = newLogger("hungry logger")

func GetLogger() *logger {
    return log
}

func (l *logger) Info(a ...interface{}) {
    fmt.Println(l.Name+":", a)
}
