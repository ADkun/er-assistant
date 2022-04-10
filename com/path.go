package com

import (
    "os"
    "path/filepath"
    "fmt"
)

// 获取当前程序工作路径
func GetWorkPath() string {
    ex, err := os.Executable()
    if err != nil {
        PanicErr(FuncName(), "os.Executable()执行失败", err)
    }

    exPath := filepath.Dir(ex)
    realPath, err := filepath.EvalSymlinks(exPath)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("filepath.EvalSymlinks(%v)执行失败", ex), err)
    }
    return realPath
}

// 更改当前工作目录
func Chdir(path string) {
    dir, err := filepath.Abs(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("filepath.Abs(%s)执行失败", path), err)
    }
    if err = os.Chdir(dir); err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Chdir(%s)执行失败", dir), err)
    }
}

func IsPathExist(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func IsDir(path string) bool {
    f, err := os.Stat(path)
    if err != nil {
        return false
    }
    return f.IsDir()
}

func IsAbs(path string) bool {
    return filepath.IsAbs(path)
}

func CreateDir(path string) {
    if IsPathExist(path) && !IsDir(path) {
        Panic(FuncName(), fmt.Sprintf("路径被占用，无法创建目录: %s", path))
    }
    if err := os.MkdirAll(path, 0777); err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.MkdirAll(%s)", path), err)
    }
    //Info(fmt.Sprintf("创建目录%s成功", path))
}

func CreateFile(path string) {
    if IsPathExist(path) && IsDir(path) {
        Panic(FuncName(), fmt.Sprintf("路径被占用，无法创建文件: %s", path))
    }

    // 创建父目录
    d := filepath.Dir(path)
    if _, err := os.Stat(d); err != nil {
        if err = os.MkdirAll(d, 0777); err != nil {
            PanicErr(FuncName(), fmt.Sprintf("os.MkdirAll(%v)", d), err)
        }
    }

    f, err := os.Create(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Create(%s)", path), err)
    }
    defer f.Close()
    //Info(fmt.Sprintf("创建文件%s成功", path))
}

