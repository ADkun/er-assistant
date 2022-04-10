package com

import (
    "os/exec"
    "bytes"
)

// 管道方式运行命令
func RunCMDPipe(command string) string {
    cmd := exec.Command(command)
    var oBuf bytes.Buffer
    cmd.Stdout = &oBuf
    if err := cmd.Start(); err != nil {
        PanicErr(FuncName(), "cmd.Start()执行失败", err)
    }
    if err := cmd.Wait(); err != nil {
        PanicErr(FuncName(), "cmd.Wait()执行失败", err)
    }
    return string(oBuf.Bytes())
}
