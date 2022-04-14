package com

import (
    "strings"
    "runtime"
)

// 作为示范，需要在相应位置调用
func GetFileLine() (string, string) {
    _, file, line, _ := runtime.Caller(0)
    return file, line
}

func Panic(funcName, msg string) {
    var builder strings.Builder
    builder.WriteString("=====[ PANIC BEG ]=====\n")
    builder.WriteString("函数: ")
    builder.WriteString(funcName)
    builder.WriteString("\n")
    builder.WriteString("信息: ")
    builder.WriteString(msg)
    builder.WriteString("\n")
    builder.WriteString("=====[ PANIC END ]=====\n")
    panic(builder.String())
}

func PanicErr(funcName, msg string, err error) {
    var builder strings.Builder
    builder.WriteString("=====[ PANIC BEG ]=====\n")
    builder.WriteString("函数: ")
    builder.WriteString(funcName)
    builder.WriteString("\n")
    builder.WriteString("信息: ")
    builder.WriteString(msg)
    builder.WriteString("\n")
    builder.WriteString("错误: ")
    builder.WriteString(err.Error())
    builder.WriteString("\n")
    builder.WriteString("=====[ PANIC END ]=====\n")
    panic(builder.String())
}

