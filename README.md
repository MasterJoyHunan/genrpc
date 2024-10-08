## genrpc 一个生成 grpc 框架的脚手架

只需定义 xx.proto 文件，genrpc 命令就可以生成整个项目的基础文件

### 使用方法

#### 安装

go 1.16 以下使用
```sh
go get -u github.com/MasterJoyHunan/genrpc
```
go 1.16 及以上使用
```sh
go install github.com/MasterJoyHunan/genrpc@v1.5.0
```

#### 初始化一个 GO 项目

```sh
mkdir you-application
cd you-application
go mod init you-app-pkg-name
```

#### 在项目下定义 xx.proto 文件

xx.proto 文件内容示例

```proto
syntax = "proto3";

package myrpc;
option go_package="proto/myrpc";


message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

service Myrpc {
  rpc Ping(Request) returns(Response);
}

```

#### 在项目下生成 grpc 项目

```sh
genrpc xx.proto
# 全部命令
genrpc --dir=. --only=server,logic --without=pb xx.proto
```

#### only 和 without 可选项

用于只生成，或不生成某个文件的代码

|参数|描述|
|-|-|
|pb|grpc 生成的 xx.pb.go 文件|
|etc|yaml 配置文件|
|config|etc 配置文件对应的 struct|
|main|主程序|
|server|grpc 服务端代码|
|logic|业务文件|
|svc|grpc 上下文|
|client|grpc 客户端代码|


#### 生成的目录结构如下

```
├─config         # 配置文件对应的 struct
├─etc            # yaml 配置文件
├─logic          # 逻辑层 - 编写业务代码的地方
├─pkg            # 
  └─grpc_client  # grpc 客户端调用文件 
├─proto          # protoc 生成的 pb 文件
├─server         # grpc 服务端文件
└─svc            # svc 文件
you-app.go       # 主程序
```

### 其他

如果觉得该项目对你有所帮助，请不要吝啬你的小手，帮忙点个 stars

如果对本项目有更好的建议或意见，欢迎提交 pr / issues，或者联系本人 tanwuyang88@gmail.com

### 协议

[MIT](https://github.com/MasterJoyHunan/genrpc/blob/master/LICENSE)
