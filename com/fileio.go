package com

import (
    "bufio"
    "os"
    "io/ioutil"
)

// 覆盖文件内容
func FWrite(path, content string) {
    if IsPathExist(path) && IsDir(path) {
        Panic(FuncName(), fmt.Sprintf("路径被占用，无法创建文件: %s", path))
    }
    f, err := os.Create(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Create(%s)", path), err)
    }
    defer f.Close()
}

// 读取文件所有内容
func FReadAll(path string) string {
    f, err := os.Open(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Open(%s)", path), err)
    }
    defer f.Close()
    return ioutil.ReadAll(f)
}
///////////////////////////////////////
// 文件读取器
// 按行读取功能
type IFReader interface {
    ReadLine() (string, bool)
}

func NewFReader(path string) *IFReader {
    file, err := os.Open(path)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.Open(%s)", path), err)
    }

    s := bufio.NewScanner(file)
    freader := &FReader {
        path: path,
        f: file,
        s: s,
    }
    return freader
}

type FReader struct {
    path string
    f *os.File
    s *bufio.Scanner
}

func (self *FReader) ReadLine() (string, bool) {
    if self.s == nil {
        Panic(FuncName(), "FReader.Scanner is nil")
    }
    end := self.s.Scan()
    line := self.s.Text()
    if err := self.s.Err(); err != nil {
        PanicErr(FuncName(), fmt.Sprintf("Scan %s failed", path), err)
    }
    line := Trim(line)
    return line, end
}

func (self *FReader) Close() {
    self.f.Close()
}

/////////////////////////////////////
// 文件写
func FAppend(path, content string) {
    // 只写模式打开文件
    f, err := os.OpenFile(path, os.O_WRONLY, 0644)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("os.OpenFile(%s)", path), err)
    }
    defer f.Close()

    // 偏移量
    n, err := f.Seek(0, os.SEEK_END)
    if err != nil {
        PanicErr(FuncName(), "f.Seek()", err)
    }
    _, err = f.WriteAt([]byte(content), n)
    if err != nil {
        PanicErr(FuncName(), "f.WriteAt()", err)
    }
}
