# MANIFEST
场景：已经生成的可执行jar包里面包含MANIFEST.MF，MANIFEST.MF里面写有引用的jar包，每次再比人的基础上加入依赖的jar包都要手动去修改一下MANIFEST.MF，
想想都觉得繁琐，一不小心少了个空格，可执行的jar包就跑不起来了。想着写个方便快捷直接生成MANIFEST.MF文件的小工具，然后就用go语言来写了它。

使用方法:使用很简单，只需要运行起来根据提示输入jar所在的目录回车即可在当前目录下生成MANIFEST.MF文件，然后编辑加入Main函数名即可。

欢迎大家指教！谢谢！
