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
		com.Pause()
	}, func() {})
}

type er struct {
	setting *com.Ini
	menu    *com.Menu
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
	c := []string{
		"1. 配置助手",
		"2. 存档管理",
		"3. 自定义备份",
		"4. MOD",
		"5. 工具",
	}
	a := []com.IAction{
		com.NewAction(self.buildMenu1()),
		com.NewAction(self.buildMenu2()),
		com.NewAction(self.buildMenu3()),
		com.NewAction(self.buildMenu4()),
		com.NewAction(self.buildMenu5()),
	}
	w := "欢迎使用艾尔登法环助手 by adkun\n"
	menu.Init(c, a, w)
	self.menu = menu
}

func (self *er) buildMenu1() *com.Menu {
	menu := com.NewMenu()
	c := []string{
		"1. 设置游戏安装目录",
		"2. 设置游戏存档目录",
	}
	one := &FuncSetGameRoot{Setting: self.setting}
	two := &FuncSetSavePath{Setting: self.setting}
	a := []com.IAction{
		com.NewAction(one),
		com.NewAction(two),
	}
	w := ""
	menu.Init(c, a, w)
	return menu
}

func (self *er) buildMenu2() *com.Menu {
	menu := com.NewMenu()
	c := []string{
		"1. 备份存档",
		"2. 恢复存档",
	}
	backup := &FuncBak{Setting: self.setting, BSave: true}
	restore := &FuncRes{Setting: self.setting, BSave: true}
	a := []com.IAction{
		com.NewAction(backup),
		com.NewAction(restore),
	}
	w := ""
	menu.Init(c, a, w)
	return menu
}

func (self *er) buildMenu3() *com.Menu {
	menu := com.NewMenu()
	c := []string{
		"1. 自定义备份",
		"2. 自定义恢复",
	}
	backup := &FuncBak{Setting: self.setting, BSave: false}
	restore := &FuncRes{Setting: self.setting, BSave: false}
	a := []com.IAction{
		com.NewAction(backup),
		com.NewAction(restore),
	}
	w := ""
	menu.Init(c, a, w)
	return menu
}

func (self *er) buildMenu4() *com.Menu {
	menu := com.NewMenu()

	list := com.ReadDir(modsDirPath)
	listLen := len(list)
	var cc []string = make([]string, listLen)
	var aa []com.IAction = make([]com.IAction, listLen)
	num := 0
	for ind, _ := range list {
		temp := modsDirPath + SLASH + list[ind]
		if !com.IsDir(temp) {
			continue
		}
		// 读取config.ini
		cfg := self.getConf(temp)
		cc[num] = self.getContent(num, cfg)
		aa[num] = com.NewAction(self.buildModMenu(cfg))
		num++
	}
	var c []string = cc[0:num]
	var a []com.IAction = aa[0:num]

	w := ""
	menu.Init(c, a, w)
	return menu
}

// Tools
func (self *er) buildMenu5() *com.Menu {
	menu := com.NewMenu()

	list := com.ReadDir(toolsDirPath)
	listLen := len(list)
	var cc []string = make([]string, listLen)
	var aa []com.IAction = make([]com.IAction, listLen)
	num := 0
	for ind, _ := range list {
		temp := toolsDirPath + SLASH + list[ind]
		if !com.IsDir(temp) {
			continue
		}
		// 读取config.ini
		cfg := self.getConf(temp)
		cc[num] = self.getContent(num, cfg)
		aa[num] = com.NewAction(self.buildModMenu(cfg))
		num++
	}
	c := cc[0:num]
	a := aa[0:num]
	w := ""
	menu.Init(c, a, w)
	return menu
}

func (self *er) buildModMenu(cfg *conf) *com.Menu {
	menu := com.NewMenu()
	c := []string{
		"1. 安装 " + cfg.sName,
        "2. 卸载 " + cfg.sName,
        "3. 查看 " + cfg.sName + " 的说明文件",
	}
	ins := &FuncIns{Setting: self.setting, Cfg: cfg}
    uni := &FuncUni{Setting:self.setting, Cfg:cfg}
    help := &FuncHelp{Cfg:cfg}
	a := []com.IAction{
		com.NewAction(ins),
        com.NewAction(uni),
        com.NewAction(help),
	}
	w := ""
	menu.Init(c, a, w)
	return menu
}

func (self *er) getContent(ind int, cfg *conf) string {
	return com.I2A(ind+1) + ". " + cfg.sName
}

func (self *er) getConf(base string) *conf {
	confIniPath := base + SLASH + cfgIniName
	confIni := com.NewIni(confIniPath)

	sName := confIni.GetString("sName")
	bCopy := confIni.GetBool("bCopy")
	bRun := confIni.GetBool("bRun")
	var sRun string
	if bRun {
		sRun = confIni.GetString("sRun")
	}
	bBak := confIni.GetBool("bBak")

	cfg := &conf{
		base:  base,
		sName: sName,
		bCopy: bCopy,
		bRun:  bRun,
		sRun:  sRun,
		bBak:  bBak,
	}
	return cfg
}
