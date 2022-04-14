package com

import (
    "strings"
)

func Panic(info []string, msg string) {
    fn := info[0]
    file := info[1]
    line := info[2]
    var builder strings.Builder
    builder.WriteString("=====[ PANIC BEG ]=====\n")
    builder.WriteString(file)
    builder.WriteString(": ")
    builder.WriteString(line)
    builder.WriteString("\n\n")
    builder.WriteString(fn)
    builder.WriteString("\n\n")
    builder.WriteString(msg)
    builder.WriteString("\n")
    builder.WriteString("=====[ PANIC END ]=====\n")
    panic(builder.String())
}

func PanicErr(info []string, msg string, err error) {
    fn := info[0]
    file := info[1]
    line := info[2]
    var builder strings.Builder
    builder.WriteString("=====[ PANICERR BEG ]=====\n")
    builder.WriteString(file)
    builder.WriteString(": ")
    builder.WriteString(line)
    builder.WriteString("\n\n")
    builder.WriteString(fn)
    builder.WriteString("\n\n")
    builder.WriteString(msg)
    builder.WriteString("\n\n")
    builder.WriteString(err.Error())
    builder.WriteString("\n")
    builder.WriteString("=====[ PANICERR END ]=====\n")
    panic(builder.String())
}

