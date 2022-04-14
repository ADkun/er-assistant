package com

import (
    "os/exec"
    "bytes"
)

// 指定工作目录，运行
func RunCd(workPath, exePath string) string {
    curPath := GetWorkPath()
    Chdir(workPath)
    res := Run(exePath)
    Chdir(curPath)
    return res
}

// 在当前工作目录下调用cmd start
func Start(exePath string) string {
    return Run("cmd", "/c", "start", exePath)
}

// 在指定工作目录下调用cmd start
func StartCd(workPath, exePath string) string {
    curPath := GetWorkPath()
    Chdir(workPath)
    res := Run("cmd", "/c", "start", exePath)
    Chdir(curPath)
    return res
}

// [IMPORTANT] 在当前工作目录下管道方式运行命令
func Run(args ...string) string {
    cmd := getCMD(args...)
    var oBuf bytes.Buffer
    cmd.Stdout = &oBuf
    if err := cmd.Start(); err != nil {
        PanicErr(DebugInfo(), "cmd.Start()执行失败", err)
    }
    if err := cmd.Wait(); err != nil {
        PanicErr(DebugInfo(), "cmd.Wait()执行失败", err)
    }
    return string(oBuf.Bytes())
}

func getCMD(args ...string) *exec.Cmd {
    ll := len(args)
    if ll == 0 {
        Panic(DebugInfo(), "参数不能为空")
    } else if ll == 1 {
        return exec.Command(args[0])
    } else if ll == 2 {
        return exec.Command(args[0], args[1])
    } else if ll == 3 {
        return exec.Command(args[0], args[1], args[2])
    } else if ll == 4 {
        return exec.Command(args[0], args[1], args[2], args[3])
    } else {
        Panic(DebugInfo(), "最大支持4个参数")
    }
    return nil
}
