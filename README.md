# 艾尔登法环助手

## 功能
1. 游戏存档备份/恢复
2. 游戏目录下自定义文件的备份/恢复
3. 可配置MOD的安装/卸载
4. 可配置工具的安装/卸载

## 说明
1. 发布版本分为“纯净版”和“打包版”
    - 纯净版：只包含助手程序本体和template/模板目录，用户需要根据该目录下的说明配置【自定义备份/MOD/Tool】
    - 打包版：包含助手程序本体和bak/、mods/、tools/，自定义备份配置了备份regulation.bin；mods/整合了一些作者自用MOD（部分只支持1.03.2）；tools/整合了一些相关工具。。
        MOD及工具均来源于网络。

## 使用方法
以“打包版”为例，假设我们要配置“梅琳娜外观MOD”，下载后发现，MOD本体只有一个文件夹parts/
说明该MOD为模型文件替换MOD。

模型文件替换MOD有两种安装方式，一种是UXM解包然后替换文件，一种是使用Mod Engine，这里以Mod Engine为例。
Mod Engine的使用方法是在游戏目录Game/下创建一个mod/文件夹，然后将模型替换mod和regulation.bin放入其中。

### 例1：梅琳娜模型替换MOD
以“梅琳娜外观MOD”为例，MOD本体是parts/xxx.dcx，配置如下（括号为说明，实际文件中不存在）：
- mods/
    - 梅琳娜外观MOD/
        - config.ini
            sName=梅琳娜外观MOD
            bCopy=true （开启复制文件到Game/选项）
            bRun=false （关闭运行选项）（该MOD没有可执行文件）
            bBak=false （关闭备份相同文件选项）
            bSkip=fales （不覆盖已存在文件）
        - files/
            - mod/
                am_m_1130.partsbnd.dcx
                am_m_1130_l.partsbnd.dcx

### 例2：反作弊切换器
“反作弊切换器”属于工具，为分类方便将器纳入“工具”范畴，因此配置在tools/下。tools/与mods/的配置方法完全相同

了解到，“反作弊切换器”需要将其文件复制到游戏目录Game/下，并运行toggle_anti_cheat.exe进行反作弊模式的切换。

配置如下：
- tools/
    - 反作弊切换器/
        - config.ini
            sName=反作弊切换器
            bCopy=true （开启复制）
            bRun=true （开启运行）
            sRun=toggle_anti_cheat.exe （运行的可执行文件“相对路径”）
            bBak=false （关闭备份）
            bSkip=true （不覆盖已存在文件）
        - files/
            - toggle_anti_cheat.exe
            - ... （此处省略n个文件或文件夹）

## 注
1. MOD和工具的config.ini中
    1. 所有"b"开头的配置项，值只能为"true"（开启）或"false"（关闭），不区分大小写。
    2. 只有当bCopy=true时，才需要配置bBak和bSkip（即当Game/下有相同文件时，备份）
    3. 只有当bRun=true时，才需要设定sRun=配置。
    4. files.ini中，FilesNum要对应下面设定的条数，条数序号从0开始。
