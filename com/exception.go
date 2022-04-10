package com

func Panic(funcName, msg string) {
    var builder string.Builder
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
    var builder string.Builder
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
