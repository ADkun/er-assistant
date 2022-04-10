package com

import (
    "io/ioutil"
    "fmt"
    "path/filepath"
    "bufio"
)

// 复制源到目标，递归复制
func Copy(from, to string) {
    f, err := os.Stat(from)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Stat(%s)", from), err)
    }
    if f.IsDir() {
        copyDir(from, to)
    } else {
        copyFile(from, to)
    }
}

func copyDir(from, to string) {
    // 创建目标目录
    if !IsPathExist(to) {
        if err := os.MkdirAll(to, 0777); err != nil {
            PanicErr(FuncName(), fmt.Sprintf("os.MkdirAll(%s)", to), err)
        }
    }

    // 复制子文件
    list, err := ioutil.ReadDir(from)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("ioutil.ReadDir(%s)", from), err)
    }
    for ind, _ := range list {
        Copy(filepath.Join(from, list[ind].Name()), filepath.Join(to, list[ind].Name()))
    }
}

func copyFile(from, to string) {
    file, err := os.Open(from)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Open(%s)", from), err)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    out, err := os.Create(to)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Create(%s)", to), err)
    }
    defer out.Close()
    if _, err = io.Copy(out, reader); err != nil {
        PanicErr(FuncName(), fmt.Sprintf("io.Copy(%v, %v)", out, reader), err)
    }
}
