package er

import (
    "erassistant/com"
    "fmt"
    "time"
)

//////////////////////////////////
// 设置游戏根目录
type FuncSetGameRoot struct {
    Setting *com.Ini
}

func (self *FuncSetGameRoot) Go() {
    for {
        com.Clear()
        fmt.Println("请输入您的游戏目录路径")
        fmt.Println("例: E:\\SteamLibrary\\steamapps\\common\\ELDEN RING\\Game")
		fmt.Println("输入(q返回):")
        path := com.ReadLine()
        if path == "q" {
            break
        }

        // 校验输入的路径
        if !self.checkPath(path) {
            time.Sleep(time.Second * 2)
            continue
        }

        // 写入配置
        self.Setting.Set(keyGameRoot, path)
        com.Info("配置成功")
        com.Pause()
        break
    }
}

func (self *FuncSetGameRoot) checkPath(path string) bool {
    if !com.IsPathExist(path) {
        com.Error("路径不存在，请重新输入")
        return false
    }

    if !com.IsDir(path) {
        com.Error("路径不是目录，请重新输入")
        return false
    }

    if !com.IsAbs(path) {
        com.Error("路径不是绝对路径，请重新输入")
        return false
    }

    // 检查是否是游戏目录
    temp := path + "/eldenring.exe"
    if !com.IsPathExist(temp) {
        com.Error("路径不是游戏目录，请重新输入（依据: 路径下无eldenring.exe)")
        return false
    }
    return true
}

//////////////////////////////////
// 设置存档目录
type FuncSetSavePath struct {
    Setting *com.Ini
}

func (self *FuncSetSavePath) Go() {
    for {
        com.Clear()
        fmt.Println("请输入您的存档目录路径")
        fmt.Println("例: C:\\Users\\Administrator\\AppData\\Roaming\\EldenRing")
		fmt.Println("输入(q返回):")
        path := com.ReadLine()
        if path == "q" {
            break
        }

        if !self.checkPath(path) {
            time.Sleep(time.Second * 2)
            continue
        }

        self.Setting.Set(keySavePath, path)
        com.Info("配置成功")
        com.Pause()
        break
    }
}

func (self *FuncSetSavePath) checkPath(path string) bool {
    if !com.IsPathExist(path) {
        com.Error("路径不存在，请重新输入")
        return false
    }

    if !com.IsDir(path) {
        com.Error("路径不是目录，请重新输入")
        return false
    }

    if !com.IsAbs(path) {
        com.Error("路径不是绝对路径，请重新输入")
        return false
    }

    temp := path + "/GraphicsConfig.xml"
    if !com.IsPathExist(temp) {
        com.Error("路径不是存档目录，请重新输入(依据: 路径下无GraphicsConfig.xml)")
        return false
    }
    return true
}

//////////////////////////////////////////////////////////
// 备份存档
type FuncBak struct {
    Setting *com.Ini // 外部
    BSave bool // 外部，true: 备份存档, false: 自定义备份
    bakDir string
    bakIni *com.Ini
    savePath string
    gamePath string
    comment string
    ts string
    tsPath string
    tarPath string

    filesNum int
    arr *com.Array
}

func (self *FuncBak) Go() {
    com.Clear()
    self.getBakDir() // save/
    if self.BSave {
        self.getKeySave() // 从配置文件获取存档目录
    } else {
        self.getKeyGamePath() // 获取游戏目录
        self.getConfig()
    }
    self.checkBakDir() // 创建备份目录，备份ini对象

    fmt.Println("请输入本次备份备注(q返回):")
    comment := com.ReadLine()
    if comment == "q" {
        return
    }
    self.comment = comment

    self.getTS() // 获取时间戳
    self.createTSDir() // 创建时间戳目录
    self.writeComment() // 创建并写入备注ini
    self.createFilesDir() // 创建save/ts/files/
    self.copyFiles() // 复制文件
    self.updateBakIni() // 更新backups.ini

    com.Info("备份成功")
    com.Pause()
}

func (self *FuncBak) updateBakIni() {
    bakNumStr := self.bakIni.Get(keyBakNum)
    if bakNumStr == "" {
        // 第一次备份
        self.bakIni.Set(keyBakNum, "1")
        maxNumStr := com.I2A(0)
        self.bakIni.Set(maxNumStr, self.ts)
    } else {
        maxNum := com.A2I(bakNumStr)
        newBakNum := maxNum + 1
        newBakNumStr := com.I2A(newBakNum)
        self.bakIni.Set(keyBakNum, newBakNumStr)
        maxNumStr := com.I2A(maxNum)
        self.bakIni.Set(maxNumStr, self.ts)
    }
}

func (self *FuncBak) copyFiles() {
    if self.BSave {
        self.copyFilesSave()
    } else {
        self.copyFilesCustom()
    }
}

func (self *FuncBak) copyFilesSave() {
    src := self.savePath
    tar := self.tarPath
    com.Copy(src, tar)
}

func (self *FuncBak) copyFilesCustom() {
    for i := 0; i < self.filesNum; i++ {
        relPath, ok := self.arr.Get(i).(string)
        if !ok {
            com.Panic(com.FuncName(), fmt.Sprint("Array.Get(%d).(string)", i))
        }
        // 检验是否存在文件
        src := self.gamePath + SLASH + relPath
        if !com.IsPathExist(src) {
            com.Panic(com.FuncName(), fmt.Sprintf("路径不存在: %s", src))
        }
        tar := self.tarPath
        tar += SLASH + relPath
        com.Copy(src, tar)
        com.Info(fmt.Sprintf("%s 备份成功", relPath))
    }
}

// 返回文件数量，文件列表
func (self *FuncBak) getConfig() {
    cfgIniPath := self.bakDir + SLASH + cfgIniName
    if !com.IsPathExist(cfgIniPath) {
        com.Panic(com.FuncName(), fmt.Sprintf("未找到自定义备注配置文件: %s", cfgIniPath))
    }

    cfgIni := com.NewIni(cfgIniPath)
    filesNumStr := cfgIni.Get(keyFilesNum)
    if filesNumStr == "" {
        com.Panic(com.FuncName(), fmt.Sprintf("无法读取FilesNum于%s，请先自行配置", cfgIniPath))
    }
    if !com.IsDigit(filesNumStr) {
        com.Panic(com.FuncName(), fmt.Sprintf("FilesNum不合法: %s", filesNumStr))
    }
    filesNum := com.A2I(filesNumStr)
    arr := com.GetArray(filesNum)
    for i := 0; i < filesNum; i++ {
        relPath := cfgIni.Get(com.I2A(i))
        if relPath == "" {
            com.Panic(com.FuncName(), fmt.Sprintf("%s 序号 %d 读取失败", cfgIniPath, i))
        }
        arr.Add(i, relPath)
    }
    self.filesNum = filesNum
    self.arr = arr
}

func (self *FuncBak) createFilesDir() {
    self.tarPath = self.tsPath + SLASH + filesDirName
    com.CreateDir(self.tarPath)
}

func (self *FuncBak) writeComment() {
    comPath := self.tsPath + SLASH + commentName // save/ts/comment.txt
    com.FWrite(comPath, self.comment)
}

func (self *FuncBak) getBakDir() {
    if self.BSave {
        self.bakDir = saveBakDirPath // save/
    } else {
        self.bakDir = bakDirPath // bak/
    }
}

func (self *FuncBak) createTSDir() {
    self.tsPath = self.bakDir + SLASH + self.ts //
}

func (self *FuncBak) getTS() {
    self.ts = com.GetCurTimeStamp()
}

func (self *FuncBak) getKeySave() {
    res := self.Setting.Get(keySavePath)
    if res == "" {
        com.Panic(com.FuncName(), "未设置存档目录，请先设置")
    }
    self.savePath = res
}

func (self *FuncBak) getKeyGamePath() {
    res := self.Setting.Get(keyGameRoot)
    if res == "" {
        com.Panic(com.FuncName(), "未设置游戏目录，请先设置")
    }
    self.gamePath = res
}

func (self *FuncBak) checkBakDir() {
    com.CreateDir(self.bakDir) // save or bak
    bakIniPath := self.bakDir + SLASH + bakIniName // xxx/backups.ini
    self.bakIni = com.NewIni(bakIniPath)
}

/////////////////////////////////////////////////////////////////
// 恢复存档
type FuncRes struct {
    Setting *com.Ini // 外部
    BSave bool // 外部

    bakIniPath string
    savePath string
    gamePath string
    bakIni *com.Ini

    bakNum int
    bakPaths *com.Array
}

func (self *FuncRes) Go() {
    if self.BSave {
        self.getKeySave()
    } else {
        self.getKeyGamePath()
    }
    self.checkBakDir()
    self.initBakIni()
    self.checkBakIni()
    ind, quit := self.selectBak() // 展示并选择备份
    if quit {
        return
    }
    self.restore(ind)
    com.Info("恢复成功")
    com.Pause()
}

func (self *FuncRes) restore(ind int) {
    tsPath, ok := self.bakPaths.Get(ind).(string)
    if !ok {
        com.Panic(com.FuncName(), fmt.Sprintf("Array.Get(%d).(string), Array:%v", ind, self.bakPaths))
    }

    src := tsPath + SLASH + filesDirName
    var tar string
    if self.BSave {
        tar = self.savePath
    } else {
        tar = self.gamePath
    }
    com.Copy(src, tar)
}

func (self *FuncRes) getKeySave() {
    self.savePath = self.Setting.Get(keySavePath)
}

func (self *FuncRes) getKeyGamePath() {
    self.gamePath = self.Setting.Get(keyGameRoot)
}

func (self *FuncRes) checkBakDir() {
    if self.BSave {
        self.bakIniPath = saveBakDirPath + SLASH + bakIniName
    } else {
        self.bakIniPath = bakDirPath + SLASH + bakIniName
    }
    if !com.IsPathExist(self.bakIniPath) {
        com.Panic(com.FuncName(), "尚未进行过备份，请先进行备份")
    }
}

func (self *FuncRes) initBakIni() {
    self.bakIni = com.NewIni(self.bakIniPath)
}

func (self *FuncRes) checkBakIni() {
    bakNumStr := self.bakIni.Get(keyBakNum)
    if bakNumStr == "" {
        com.Panic(com.FuncName(), "尚未进行过备份，请先进行备份")
    }
    self.bakNum = com.A2I(bakNumStr)
    self.bakPaths = com.GetArray(self.bakNum)
    for i := 0; i < self.bakNum; i++ {
        k := com.I2A(i)
        ts := self.bakIni.Get(k)
        if ts == "" {
            com.Panic(com.FuncName(), fmt.Sprintf("backups.ini文件损坏，序号 %d 不存在", i))
        }
        var tsPath string
        if self.BSave {
            tsPath = saveBakDirPath + SLASH + ts
        } else {
            tsPath = bakDirPath + SLASH + ts
        }
        if !com.IsPathExist(tsPath) {
            com.Panic(com.FuncName(), fmt.Sprintf("backups.ini对应的备份文件丢失，序号 %d", i))
        }
        // 获取并存储备份的ts路径 save/ts
        self.bakPaths.Add(i, tsPath)
    }
}

// 选择备份序号
func (self *FuncRes) selectBak() (int, bool) {
    for {
        com.Clear()
        // 打印
        fmt.Println("已备份: ")
        for i := 0; i < self.bakNum; i++ {
            tsPath, ok := self.bakPaths.Get(i).(string)
            if !ok {
                com.Panic(com.FuncName(), fmt.Sprintf("Array.Get(%d).(string), Array:%v", i, self.bakPaths))
            }
            // 获取备注
            commentPath := tsPath + SLASH + commentName
            if !com.IsPathExist(commentPath) {
                com.Panic(com.FuncName(), fmt.Sprintf("未找到备注文件: %s", commentPath))
            }
            comment := com.FReadAll(commentPath)
            pStr := com.I2A(i + 1) + ": " + comment
            fmt.Println(pStr)
        }

        // 选择
        fmt.Println()
        fmt.Println("选择要恢复的存档序号(q返回): ")
        inp := com.ReadLine()
        if inp == "q" {
            return -1, true // 返回
        }
        if !com.IsDigit(inp) {
            com.Error("输入不合法，请重新输入")
            com.Pause()
            continue
        }
        ind := com.A2I(inp)
        if ind <= 0 || ind > self.bakNum {
            com.Error("输入范围不合法")
            com.Pause()
            continue
        }
        return ind - 1, false
    }
}

////////////////////////////////////////////
// 安装MOD/工具
type FuncIns struct {
    Setting *com.Ini
    Cfg *conf

    gamePath string
    basePath string
    filesIniPath string
    filesDirPath string
    filesIni *com.Ini
    filesNum int
    filesRelPath []string
    bakPath string
}

func (self *FuncIns) Go() {
    self.showCfg()
    self.initParams()
    self.backup()
    self.copy()
    self.run()
    com.Info("ok")
    com.Pause()
}

func (self *FuncIns) showCfg() {
    if self.Cfg.bBak {
        com.Info("备份: 是")
    } else {
        com.Info("备份: 否")
    }

    if self.Cfg.bRun {
        com.Info("运行: 是")
        com.Info("运行程序: " + self.Cfg.sRun)
    } else {
        com.Info("运行: 否")
    }

    if self.Cfg.bCopy {
        com.Info("复制: 是")
    } else {
        com.Info("复制: 否")
    }
    fmt.Println()
}

func (self *FuncIns) backup() {
    if !self.Cfg.bBak {
        return
    }

    com.Info("开始备份")
    for ind, _ := range self.filesRelPath {
        src := self.gamePath + SLASH + self.filesRelPath[ind]
        if !com.IsPathExist(src) {
            com.Warn(fmt.Sprintf("%s 不存在，跳过备份", src))
            continue
        }

        tar := self.bakPath + SLASH + self.filesRelPath[ind]
        if com.IsPathExist(tar) {
            com.Warn(fmt.Sprintf("%s 已存在，跳过备份", tar))
        }
        com.Copy(src, tar)
        com.Info(fmt.Sprintf("已备份 %s", self.filesRelPath[ind]))
    }
    com.Info("备份完成")
}

func (self *FuncIns) copy() {
    if !self.Cfg.bCopy {
        return
    }

    com.Info("开始复制")
    for ind, _ := range self.filesRelPath {
        relPath := self.filesRelPath[ind]
        src := self.filesDirPath + SLASH + relPath
        if !com.IsPathExist(src) {
            com.Panic(com.FuncName(), fmt.Sprintf("%s 不存在", src))
        }
        tar := self.gamePath + SLASH + relPath
        com.Copy(src, tar)
        com.Info("复制 " + relPath)
    }
    com.Info("复制完毕")
}

func (self *FuncIns) run() {
    if !self.Cfg.bRun {
        return
    }

    com.Info("开始运行")
    exePath := self.getExePath()
    newWorkPath := self.getNewWorkPath()
    com.RunCd(newWorkPath, exePath)
    //com.Info("运行输出:\n" + output)
    com.Info("运行结束")
}

func (self *FuncIns) getNewWorkPath() string {
    if !self.Cfg.bRun {
        return ""
    }

    if self.Cfg.bCopy {
        res := com.Dir(self.gamePath + SLASH + self.Cfg.sRun)
        return res
    } else {
        res := com.Dir(self.filesDirPath + SLASH + self.Cfg.sRun)
        return res
    }
    return ""
}

func (self *FuncIns) getExePath() string {
    if !self.Cfg.bRun {
        return ""
    }

    if self.Cfg.bCopy {
        return self.gamePath + SLASH + self.Cfg.sRun
    } else {
        return self.Cfg.sRun
    }
    return ""
}

func (self *FuncIns) initParams() {
    self.gamePath = self.getGamePath() // 游戏目录
    self.basePath = self.Cfg.base // mod/tool目录
    self.filesIniPath = self.basePath + SLASH + filesIniName // files.ini路径
    self.filesDirPath = self.basePath + SLASH + "files"
    self.filesIni = com.NewIni(self.filesIniPath)
    self.filesNum, self.filesRelPath = self.readFilesIni()
    self.bakPath = self.basePath + SLASH + "bak"
}

func (self *FuncIns) getBakPath() string {
    return self.basePath + "SLASH" + "bak"
}

func (self *FuncIns) readFilesIni() (int, []string) {
    filesNum := self.filesIni.GetInt(keyFilesNum)
    filesRelPath := make([]string, filesNum)
    for i := 0; i < filesNum; i++ {
        relPath := self.filesIni.GetString(com.I2A(i))
        filesRelPath[i] = relPath
    }
    return filesNum, filesRelPath
}

func (self *FuncIns) getGamePath() string {
    return self.Setting.GetString(keyGameRoot)
}

///////////////////////////////////////////
// 卸载MOD/工具
type FuncUni struct {
    Setting *com.Ini
    Cfg *conf

    gamePath string
    basePath string
    filesIniPath string
    filesIni *com.Ini
    filesNum int
    filesRelPath []string
    bakDirPath string
}

func (self *FuncUni) Go() {
    self.getParams()
    self.rmFiles()
    self.restore()
    com.Info("ok")
    com.Pause()
}

func (self *FuncUni) restore() {
    if !self.Cfg.bCopy || !self.Cfg.bBak {
        return
    }

    com.Info("开始恢复")
    for ind, _ := range self.filesRelPath {
        relPath := self.filesRelPath[ind]
        src := self.bakDirPath + SLASH + relPath
        if !com.IsPathExist(src) {
            com.Info(fmt.Sprintf("%s 不存在，跳过恢复", src))
            continue
        }
        tar := self.gamePath + SLASH + relPath
        com.Copy(src, tar)
        com.Info("恢复 " + relPath)
    }
    com.Info("恢复完毕")
}

func (self *FuncUni) rmFiles() {
    if !self.Cfg.bCopy {
        return
    }

    com.Info("开始删除")
    for ind, _ := range self.filesRelPath {
        relPath := self.filesRelPath[ind]
        temp := self.gamePath + SLASH + relPath
        com.RemoveAll(temp)
        com.Info("删除文件 " + temp)
    }
    com.Info("删除完毕")
}

func (self *FuncUni) getParams() {
    self.gamePath = self.Setting.GetString(keyGameRoot)
    self.basePath = self.Cfg.base
    self.filesIniPath = self.basePath + SLASH + filesIniName
    self.filesIni = com.NewIni(self.filesIniPath)
    self.filesNum = self.filesIni.GetInt("FilesNum")
    filesRelPath := make([]string, self.filesNum)
    for i := 0; i < self.filesNum; i++ {
        relPath := self.filesIni.GetString(com.I2A(i))
        filesRelPath[i] = relPath
    }
    self.filesRelPath = filesRelPath
    self.bakDirPath = self.basePath + SLASH + "bak"
}

////////////////////////////////////////////
// 帮助
type FuncHelp struct {
    Cfg *conf
}

func (self *FuncHelp) Go() {
    com.Clear()
    helpPath := self.Cfg.base + SLASH + "readme.txt"
    if !com.IsPathExist(helpPath) {
        com.Panic(com.FuncName(), fmt.Sprintf("帮助文件 %s 不存在", helpPath))
    }

    c := com.FReadAll(helpPath)
    fmt.Println(c)
    fmt.Println()
    com.Pause()
}
