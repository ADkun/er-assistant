package com

import (
    "os"
    "fmt"
    "path/filepath"
    "strings"
)

type IIni interface {
    Get(k string) string
    Set(k, v string)
}

func NewIni(path string) *Ini {
    // 创建父目录
    d := filepath.Dir(path)
    if _, err := os.Stat(d); err != nil {
        if err = os.MkdirAll(d, 0777); err != nil {
            PanicErr(DebugInfo(), fmt.Sprintf("os.MkdirAll(%v)", d), err)
        }
    }

    // 文件不存在则创建
    exist := IsPathExist(path)
    if !exist {
        file, err := os.Create(path)
        if err != nil {
            PanicErr(DebugInfo(), fmt.Sprintf("os.Create(%s)", path), err)
        }
        defer file.Close()
    }

    // 已存在且是目录，报错
    if IsDir(path) {
        Panic(DebugInfo(), fmt.Sprintf("路径被占用，无法创建文件%s", path))
    }

    ini := &Ini {
        path: path,
    }
    return ini
}

type Ini struct {
    path string
}

// 获取指定k对应的v
func (self *Ini) Get(k string) string {
    freader := NewFReader(self.path)
    defer freader.Close()

    for {
        r, remain := freader.ReadLine()
        if !remain {
            break
        }
        // 跳过空行
        if r == "" {
            continue
        }

        rs := strings.Split(r, "=")
        rsLen := len(rs)
        // 跳过格式错误的行
        if rsLen != 2 {
            continue
        }
        if rs[0] == k {
            res := rs[1]
            res = Trim(res)
            return res
        }
    }
    // 返回空，代表未设置值
    return ""
}

// 设置kv
func (self *Ini) Set(k, v string) {
    var builder strings.Builder
    freader := NewFReader(self.path)

    // 创建临时文件
    temp := GetRandStr(10)
    tempF, err := os.Create(temp)
    if err != nil {
        PanicErr(DebugInfo(), fmt.Sprintf("os.Create(%s)", temp), err)
    }

    // 构造
    builder.WriteString(k)
    builder.WriteString("=")
    builder.WriteString(v)
    builder.WriteString("\n")
    target := builder.String()

    var bf bool = false
    for {
        r, remain := freader.ReadLine()
        if !remain {
            break
        }
        rs := strings.Split(r, "=")
        rsLen := len(rs)
        if rsLen != 2 {
            continue
        }
        // 若是目标key，写新值
        if rs[0] == k {
            tempF.WriteString(target)
            bf = true
        } else {
            tempF.WriteString(r + "\n")
        }
    }

    if bf == false {
        tempF.WriteString(target)
    }

    freader.Close()
    if err = os.Remove(self.path); err != nil {
        PanicErr(DebugInfo(), fmt.Sprintf("os.Remove(%s)", self.path), err)
    }

    tempF.Close()
    // 临时文件覆盖原文件
    if err = os.Rename(temp, self.path); err != nil {
        PanicErr(DebugInfo(), fmt.Sprintf("os.Rename(%s)", temp), err)
    }
}

func (self *Ini) GetBool(k string) bool {
    v := self.Get(k)
    v = ToUpper(v)
    if v == "" {
        Panic(DebugInfo(), fmt.Sprintf("%s 未设置 %s", self.path, k))
    }

    if v == "TRUE" {
        return true
    } else if v == "FALSE" {
        return false
    } else {
        Panic(DebugInfo(), fmt.Sprintf("%s: %s 的值不合法: %s\n应为true或false", self.path, k, v))
    }
    return false
}

func (self *Ini) GetString(k string) string {
    v := self.Get(k)
    if v == "" {
        Panic(DebugInfo(), fmt.Sprintf("%s 未设置 %s", self.path, k))
    }
    return v
}

func (self *Ini) GetInt(k string) int {
    v := self.Get(k)
    if v == "" {
        Panic(DebugInfo(), fmt.Sprintf("%s 未设置 %s", self.path, k))
    }
    return A2I(v)
}
