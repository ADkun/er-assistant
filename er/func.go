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
type FuncBakSave struct {
    Setting *com.Ini // 外部
    bakDir string
    bakIni *com.Ini
    savePath string
    comment string
    ts string
    tsPath string
    tarPath string
}

func (self *FuncBakSave) Go() {
    com.Clear()
    self.getKeySave() // 从配置文件获取存档目录
    self.getBakDir() // save/
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

func (self *FuncBakSave) updateBakIni() {
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

func (self *FuncBakSave) copyFiles() {
    src := self.savePath
    tar := self.tarPath
    com.Copy(src, tar)
}

func (self *FuncBakSave) createFilesDir() {
    self.tarPath = self.tsPath + SLASH + filesDirName
    com.CreateDir(self.tarPath)
}

func (self *FuncBakSave) writeComment() {
    comPath := self.tsPath + SLASH + commentName // save/ts/comment.txt
    com.FWrite(comPath, self.comment)
}

func (self *FuncBakSave) getBakDir() {
    self.bakDir = saveBakDirPath // save
}

func (self *FuncBakSave) createTSDir() {
    self.tsPath = self.bakDir + SLASH + self.ts //
}

func (self *FuncBakSave) getTS() {
    self.ts = com.GetCurTimeStamp()
}

func (self *FuncBakSave) getKeySave() {
    res := self.Setting.Get(keySavePath)
    if res == "" {
        com.Panic(com.FuncName(), "未设置存档目录，请先设置")
    }
    self.savePath = res
}

func (self *FuncBakSave) checkBakDir() {
    com.CreateDir(self.bakDir) // save or bak
    bakIniPath := self.bakDir + SLASH + bakIniName // xxx/backups.ini
    self.bakIni = com.NewIni(bakIniPath)
}
