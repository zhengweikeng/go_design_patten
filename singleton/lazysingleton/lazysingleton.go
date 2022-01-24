/**
 * @Author: Seedzheng
 * 单例懒汉模式
 */

package lazysingleton

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
var m sync.Mutex

func GetLogger() *logger {
    m.Lock()
    if log == nil {
        log = newLogger("lazy logger")
    }
    m.Unlock()

    return log
}

func (l *logger) Info(a ...interface{}) {
    fmt.Println(l.Name+":", a)
}
