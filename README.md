# thrift_demo
thrift demo

### 安装
```bash
# 此处为mac安装，其余环境略
brew install thrift
```

### 使用指南(golang)
```bash
thrift --gen go hello.thrift
go mod init demo
go mod tidy
go run server.go
go run client.go
```
