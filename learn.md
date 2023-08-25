
# d1 2023-08-23 启动项目
> 遇到的坑  bee 工具需要安装，设置环境变量。
> 
> 


# d2 踩坑记录

> go env -w GO111MODULE=on 

* 知道了 golang 环境  有go mod 和 gopath 两种模块加载方式。


# day 3 踩坑记录

> 写了一个用户注册的接口
>
> orm的一些问题。
> go 没有try catch

在go 开发中喜欢使用如下的编程范式

错误信息可读性比较差。
```go
err, _ = XXX(xxx,xxx)

```
