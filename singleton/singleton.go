/**
 * @Author: Seedzheng
 * 单例模式
 */

package singleton

import (
    "fmt"
    "sync"
)

type logger struct {
    Name string
}

func newLogger(name string) *logger {
    // ... 初始化logger
    return &logger{name}
}

var log *logger
var once sync.Once

func GetLogger() *logger {
    once.Do(func() {
        log = newLogger("singleton logger")
    })

    return log
}

func (l *logger) Info(a ...interface{}) {
    fmt.Println(l.Name+":", a)
}
