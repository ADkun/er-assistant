package er

import (
    "runtime"
)

// 基础目录名
const toolsDirPath   = "tools"
const modsDirPath    = "mods"
const saveBakDirPath = "save"
const bakDirPath     = "bak"

// 程序配置文件路径
const settingIniPath = "config/settings.ini"

// 单独文件名
const filesIniName  = "files.ini" // MOD和工具的本体文件名
const modCfgIniName = "config.ini" // MOD和工具的配置文件名
const bakIniName = "backups.ini" // 备份文件记录（自动生成）
const bakCfgIniName = "config.ini" // 自定义备份配置文件
const commentName = "comment.txt"

// 单独文件夹
const filesDirName = "files"

// settings.ini 配置项
const keyGameRoot = "GameRoot" // 游戏安装根目录
const keySavePath = "SavePath" // 游戏存档根目录

// files.ini 配置项
const keyFilesNum = "FilesNum" // MOD的文件数量

// backups.ini配置项
const keyBakNum = "BakNum"

var SLASH string
func init() {
    if runtime.GOOS == "linux" {
        SLASH = "/"
    } else {
        SLASH = "\\"
    }
}
