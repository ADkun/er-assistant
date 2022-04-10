package com

import (
    "fmt"
    "time"
)

// IMenu
type IMenu interface {
    Go() // 菜单作为程序最上层
    Initialize() // 初始化
    AddContent(i int, c interface{}) // string
    AddAction(i int, a interface{}) // IAction
    monitor() // 监控用户输入
    print() // 打印菜单内容
}

func NewMenu() *Menu {
    return &Menu{}
}

type Menu struct {
    welcome string
    content *Array
    action *Array
}

func (self *Menu) Init(c []string, a []IAction, w string) {
    cLen := len(c)
    aLen := len(a)
    if cLen != aLen {
        Panic(FuncName(), fmt.Sprintf("参数数组大小不相等 %d != %d", cLen, aLen))
    }
    self.Initialize(cLen)
    self.SetWelcome(w)
    for i, _ := range c {
        self.AddContent(i, c[i])
        self.AddAction(i, a[i])
    }
}

func (self *Menu) Initialize(i int) {
    self.content = GetArray(i)
    self.action = GetArray(i)
}

func (self *Menu) SetWelcome(s string) {
    self.welcome = s
}

func (self *Menu) Go() {
    self.monitor()
}

func (self *Menu) print() {
    Clear()
    if self.welcome != "" {
        fmt.Println(self.welcome)
    }
    size := self.content.GetSize()
    for i := 0; i < size; i++ {
        fmt.Println(self.content.Get(i))
    }
}

func (self *Menu) monitor() {
    for {
        self.print()
        fmt.Println("\n请选择序号(q返回):")
        r := ReadLine()
        if r == "q" {
            break
        }


        var i int
        var berr bool = false
        Try(func() {
            i = A2I(r)
        }, func(err interface{}) {
            berr = true
        }, func(){})
        if berr || !IsDigit(r) {
            Error("输入不合法: " + r)
            time.Sleep(time.Second * 2)
            continue
        }
        if i <= 0 || i > self.action.GetSize() {
            Error("输入范围不合法: " + r)
            time.Sleep(time.Second * 2)
            continue
        }

        if v, ok := self.action.Get(i - 1).(IAction); ok {
            // Menu是最上一级接收异常的
            Try(func() {
                v.Go()
            }, func(err interface{}) {
                fmt.Println(err)
                Pause()
            }, func(){})
        } else {
            Panic(FuncName(), "action无法转换为IAction")
        }
    }
}

func (self *Menu) AddContent(i int, c interface{}) {
    cap := self.content.GetCapacity()
    if i < 0 || i > cap {
        Panic(FuncName(), fmt.Sprintf("索引错误, 最大%d", cap))
    }

    if _, ok := c.(string); !ok {
        Panic(FuncName(), "参数必须为string类型")
    }

    self.content.Add(i, c)
}

func (self *Menu) AddAction(i int, a interface{}) {
    cap := self.action.GetCapacity()
    if i < 0 || i > cap {
        Panic(FuncName(), fmt.Sprintf("索引错误, 最大%d", cap))
    }

    if _, ok := a.(IAction); !ok {
        Panic(FuncName(), "参数必须为IAction类型")
    }

    self.action.Add(i, a)
}

// IAction
type IAction interface {
    Go()
}

func NewAction(a interface{}) *Action {
    _, ok := a.(IMenu)
    _, ok1 := a.(IFunc)
    if !ok && !ok1 {
        Panic(FuncName(), "参数应为IMenu或IFunc类型")
    }
    action := &Action {
        action: a,
    }
    return action
}

type Action struct {
    action interface{} // IMenu or IFunc
}

func (self *Action) Go() {
    if v, ok := self.action.(IMenu); ok {
        v.Go()
    }

    if v, ok := self.action.(IFunc); ok {
        v.Go()
    }
}

// IFunc
type IFunc interface {
    Go() // 子类自定义
}

