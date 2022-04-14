package com

import (
    "os"
    "fmt"
)

func GetArgLen() int {
    return len(os.Args)
}

func GetArgs() []string {
    return os.Args
}

func GetArgAt(i int) string {
    l := GetArgLen()
    if i > l - 1 {
        Panic(DebugInfo(), fmt.Sprintf("索引超出范围，索引范围: 0 - %d", l - 1))
    }
    args := GetArgs()
    return args[i]
}

// 自定义参数
type NamedFlag struct {

}

func GetNamedFlag() *NamedFlag {
    return nil
}
