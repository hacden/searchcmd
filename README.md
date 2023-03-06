命令搜索工具，增删改查，通过编辑yaml文件添加到数据库文件中保存和最后检索，快速搜索日常渗透测试工作中使用到的工具命令。

### searchcmd
```
1、编译window
go env -w GO111MODULE=auto
go env -w GOOS=windows
go env -w GOARCH=386
go build -ldflags="-s -w" -trimpath main.go

4、编译linux
go env -w GO111MODULE=auto
go env -w GOOS=linux
go env -w GOARCH=386
go build -ldflags="-s -w" -trimpath main.go

```

### 简单示例
1. 关键字搜索
```
searchcmd -name useradd
```
2. 模糊搜索
```
searchcmd -like -name shell
```
3. 更新索引文件(前提是需要yaml文件)
```
searchcmd -update
```
4. 插入索引文件(前提是需要yaml文件)
```
searchcmd -insert
```
5. 删除条目
```
searchcmd -delete -name example
```
6. 列出所有支持搜索的模块
```
searchcmd -all
```
7. 添加或者更新，请参考example.yaml进行编写
```
./example.yaml
```


### 参考优秀项目: 
感谢 Rvn0xsy [red-tldr](https://github.com/Rvn0xsy/red-tldr)