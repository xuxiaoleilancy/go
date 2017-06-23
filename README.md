# go 工程目录

## rpackage 
包含了对go基础库的扩展，例如osextend，stringextend等扩展

### 在osextend中你可以使用以下接口

```
GetFiles 获取指定目录下指定类型的文件列表
CopyFile 拷贝文件
ReadLine 按行读取文件内容
PathExists 检查给定目录是否存在
......
```

## r_build_tools 
r_build_tools 包含了针对随锐R框架开发的一些小工具
