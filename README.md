# 规划实施工具InfraRobot

本工具可以将Excel中的主机规划内容，转化成在主机上实施的脚本。


## 目录结构

发布包解压缩后infra-robot中包含以下内容:

 * config/dict.json - 应用的字典，其中列出了所有的关键字及其别名
 * config/tmpl - 默认的模板文件
 * sample - 包含了示例规划文件
 * infra-robot - MAC下的可执行文件
 * infra-robot.exe - Windows下的可执行文件

## 安装

将infra-robot目录加入到系统的PATH中，可以直接在命令行里执行infra-robot.exe或infra-robot即可。

## 执行

将规划Excel存放到一个目录，然后执行:

```
infra-robot targetdirpath
```

这时会直接使用config/tmpl目录下的模板，如果想要使用自己的模板，可以如下方式执行：

```
infra-robot targetdirpath tmpldirpath
```

## 定制字典

config/dict.json包含了所有字典，可以根据需要添加调整。host，hosts，module等是特殊字典，一般不要更改。

## 模板编辑

模板使用的是go语言模板语言，请参考相关文档学习复杂的编辑。

