[ 说明 ]
mods/
自定义配置的Mod目录，将Mod文件按指定路径放入mods/下的独立文件夹中，助手会自动读取配置并执行相应的行为，对MOD进行安装/卸载/运行

与tools/的配置方法相同，仅作目录以及程序逻辑上的区分

[ 文件结构 ]
- MOD目录名/         任意MOD目录名
    - config.ini     MOD安装/卸载相关配置
    - files.ini      指定MOD文件数量，以及相对于游戏根目录Game/下的相对路径
    - files/         存放具体的MOD文件，文件路径需要符号Game/下的相对路径，并配置到files.ini中

[ config.ini配置项 ]
- 
