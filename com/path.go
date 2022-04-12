package com

import (
    "os"
    "path/filepath"
    "fmt"
    "io/ioutil"
)

func Dir(path string) string {
    return filepath.Dir(path)
}

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

func ReadDir(path string) []string {
    if !IsPathExist(path) {
        Panic(FuncName(), fmt.Sprintf("路径不存在: %s", path))
    }

    list, err := ioutil.ReadDir(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("ioutil.ReadDir(%s)", path), err)
    }

    var res []string = make([]string, len(list))
    for i := 0; i < len(list); i++ {
        res[i] = list[i].Name()
    }
    return res
}

func RemoveAll(path string) {
    if err := os.RemoveAll(path); err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.RemoveAll(%s)", path), err)
    }
}

func Abs(path string) string {
    res, err := filepath.Abs(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("filepath.Abs(%s)", path), err)
    }
    return res
}

// 读取一个目录下所有的
func ReadAllFiles(path string) []string {
    if !IsDir(path) {
        Panic(FuncName(), fmt.Sprintf("%s 不是目录", path))
    }

    // 读取第一层级目录
    arr := GetArray(50)
    readAllFiles(path, "", arr)

    // 转换为[]string
    s := arr.GetSize()
    res := make([]string, s)
    for i := 0; i < s; i++ {
        if v, ok := arr.Get(i).(string); ok {
            res[i] = v
        } else {
            Panic(FuncName(), fmt.Sprintf("Array元素转换为string失败"))
        }
    }

    return res
}

// var count int = 1

func readAllFiles(path string, prefix string, arr *Array) {
    list := ReadDir(path)
    for ind, _ := range list {
        name := list[ind]
        tPath := path + SLASH + name

        if IsDir(tPath) {
            if prefix == "" {
                readAllFiles(tPath, name, arr)
            } else {
                readAllFiles(tPath, prefix + SLASH + name, arr)
            }
            continue
        }

        if prefix == "" {
            arr.AddLast(name)
        } else {
            a := prefix + SLASH + name
            arr.AddLast(a)
        }
    }
}
