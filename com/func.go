package com

import (
    "runtime"
    "path/filepath"
)

func DebugInfo() []string {
    pc := make([]uintptr, 1)
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    funcName := f.Name()
    file, line := f.FileLine(pc[0])
    ll := I2A(line)
    baseFile := filepath.Base(file)
    return []string{funcName, baseFile, ll}
}

func FileLine() (string, int) {
    pc := make([]uintptr, 1)
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    return f.FileLine(pc[0])
}

func FuncName() string {
    pc := make([]uintptr, 1)
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    return f.Name()
}
