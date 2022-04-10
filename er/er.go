package er

import (
    "erassistant/com"
    "fmt"
)

func Entry() {
    com.Try(func() {

        er := &er{}
        er.initIni()
        er.buildMenu()
        er.showMenu()

    }, func(err interface{}) {
        fmt.Println(err)
    }, func(){})
}

type er struct {
    setting *com.Ini
    menu *com.Menu
}

func (self *er) initIni() {
    self.setting = com.NewIni(settingIniPath) // config/settings.ini
}

func (self *er) showMenu() {
    self.menu.Go()
}

func (self *er) buildMenu() {
    // 根目录
    menu := com.NewMenu()
    c := []string {
        "1. 配置助手",
        "2. 存档管理",
        //"3. 自定义备份",
        //"4. MOD",
        //"5. 工具",
    }
    a := []com.IAction {
        com.NewAction(self.buildMenu1()),
        com.NewAction(self.buildMenu2()),
        //NewAction(self.buildMenu3()),
        //NewAction(self.buildMenu4()),
        //NewAction(self.buildMenu5()),
    }
    w := "欢迎使用艾尔登法环助手 by adkun\n"
    menu.Init(c, a, w)
    self.menu = menu
}

func (self *er) buildMenu1() *com.Menu {
    menu := com.NewMenu()
    c := []string {
        "1. 设置游戏安装目录",
        "2. 设置游戏存档目录",
    }
    one := &FuncSetGameRoot{Setting:self.setting}
    two := &FuncSetSavePath{Setting:self.setting}
    a := []com.IAction {
        com.NewAction(one),
        com.NewAction(two),
    }
    w := ""
    menu.Init(c, a, w)
    return menu
}

func (self *er) buildMenu2() *com.Menu {
    menu := com.NewMenu()
    c := []string {
        "1. 备份存档",
        //"2. 恢复存档",
    }
    backup := &FuncBakSave{Setting:self.setting}
    //restore := &FuncResSave{setting:self.setting}
    a := []com.IAction {
        com.NewAction(backup),
        //com.NewAction(restore),
    }
    w := ""
    menu.Init(c, a, w)
    return menu
}
