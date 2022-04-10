package com

import (
    "bufio"
    "os/exec"
    "fmt"
)

func ReadLine() string {
    reader := bufio.NewReader(os.Stdin)
    res, err := reader.ReadString('\n')
    if err != nil {
        PanicErr(FuncName(), "reader.ReadString('\n')", err)
    }
    return Trim(res)
}

func Pause() {
    fmt.Println("按回车键继续...")
    fmt.Scanf("%s")
}

func Info(s string) {
    fmt.Println("[ INFO ] " + s)
}

func Warn(s string) {
    fmt.Println("[ WARN ] " + s)
}

func Error(s string) {
    fmt.Println("[ ERROR ] " + s)
}

func Clear() {
    cmd := exec.Command("cmd.exe", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
