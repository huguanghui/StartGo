# StartGo

开始接触GO的入门练习.

## Go的历史

Google在2009年发布的编程语言.
罗布·派克(Rob Pike)
Thompson
Robert Griesemer

在不损失应用程序性能的情况下降低代码的复杂性

## type

### struct的使用

```go
type tcpKeepAliveListener struct {
    *net.TCPListener
}
```

## Net

### tcp

```go
net.Listen("tcp", addr)
l.Accept()
```