package com

import (
    "os"
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
        Panic(FuncName(), fmt.Sprintf("索引超出范围，索引范围: 0 - %d", len - 1))
    }
    args := GetArgs()
    return args[i], nil
}

// 自定义参数
type NamedFlag struct {

}

func GetNamedFlag() *NamedFlag {
    return nil
}
