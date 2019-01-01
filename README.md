# StartGo

开始接触GO的入门练习.

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