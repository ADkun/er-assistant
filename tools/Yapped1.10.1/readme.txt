Yapped-Rune-Bear v1.07

[ 地址 ]
https://github.com/vawser/Yapped-Rune-Bear

[ 说明 ]
Yapped是一款用来编辑.bin文件的工具，在艾尔登法环中，对应regulation.bin

[ 使用方法 ]
1. 下载MOD的.csv文件
2. 将.csv文件放到Yapped目录下的res/GR/Data目录下（使用本助手安装后，Yapped位于游戏目录Game/YAPPED
3. 运行Yapped
4. 如果是第一次运行，那么
    1. 点击File->Dark Soul 3（下拉选择Elden Ring）
    2. 点击File->Open，进入到游戏Game目录下，选择regulation.bin
    3. 手动备份一份未经修改的regulation.bin
5. 以隐形头盔MOD为例，下载下来的是"EquipParamProtector.csv"，
    1. 在Yapped界面显示的三栏中的最左边一栏找到"EquipParamProtector"，单击选中它
    2. 点击Tools->Import Data->是(Y)
    3. 点击File->Save，即可更新regulation.bin

[ 注 ]
由于Yapped是外部程序，不支持自动MOD替换，故将一些需要Yapped工具的MOD放置在助手目录：
mod/版本号/yapped/mod/下
每个MOD都有对应的readme.txt说明。

[ 导入错误的解决方法 ]
参考：https://www.bilibili.com/video/BV1NS4y1m7iT
如果导入csv文件的时候，报错，说明分隔符不对。
将Yapped的分隔符改为英文分号，再用WPS打开.csv文件，选择“数据->分列->下一步->勾选'分号'->下一步->保存”，然后保存，再导入即可。