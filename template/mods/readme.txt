自定义存放MOD的文件夹

[ 说明 ]
该目录下每个子目录，都视作一个MOD，子目录结构必须按照要求配置，
子目录名随意，方便辨识即可。

[ 目录结构 ]
- mods/
    - 示例MOD文件夹/：该目录视作一个独立MOD的文件夹
        - files/：MOD本体文件目录，对应于游戏目录Game/
        - config.ini：对于MOD安装/卸载相关的配置
            - sName=示例MOD：表示MOD名称为“示例MOD”，将会显示在助手界面上。
            - bCopy=true：表示开启复制选项，将会将files.ini中指定的文件从files/目录下复制到游戏目录Game/
            - bRun=true：表示开启运行选项，将会运行sRun参数指定的程序。
            - sRun=MOD主程序.exe：表示运行的程序相对路径，如果程序位于子目录，可以写成“XXX/MOD主程序.exe”（仅当bRun=true时需要设置）
            - bBak=false：表示不开启备份选项。如果开启备份，当bCopy=true，且Game/目录下有对应的文件时，安装MOD时会将其备份；卸载MOD时会将其备份。

[ 注 ]
1. config.ini中"b"开头的参数值可以是"true"（开启），或者"false"（关闭），不区分大小写。