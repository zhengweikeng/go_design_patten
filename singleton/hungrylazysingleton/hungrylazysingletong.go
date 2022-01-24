/**
 * @Author: Seedzheng
 * 单例饿汉&懒汉结合
 */

package hungrylazysingleton

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
    if log == nil {
        m.Lock()
        if log == nil {
            log = newLogger("hungry&lazy logger")
        }
        m.Unlock()
    }

    return log
}

func (l *logger) Info(a ...interface{}) {
    fmt.Println(l.Name+":", a)
}
