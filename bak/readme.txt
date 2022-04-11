[ 说明 ]
自定义备份文件夹
路径: bak/

[ 文件描述 ]
backups.ini: 备份时自动生成的记录文件，请勿随意修改
config.ini: 自定义配置文件

[ config.ini配置项 ]
- FilesNum=2
    表示一共需要备份2个文件（或文件夹），对应下面的项
- 0=regulation.bin
    表示需要备份游戏根目录Game/下的regulation.bin（相对路径）
- 1=mod/
    表示需要备份游戏根目录Game/下的mod文件夹（"/"可有可无）

[ 注意事项 ]
文件末尾需要保留一个空行
