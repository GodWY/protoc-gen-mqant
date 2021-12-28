# protoc-gen-mqant
基于mqant框架的rpc协议自动生成代码工具

# 自动生成协议

## 工具
1. 下载自动化生成工具
> git clone https://tygit.tuyoo.com/dev5/protoc-gen-mqant.git

2. 编译文件

```go
    go build -o protoc-gen-mqant main.go
```
3. 将protoc-gen-mqant加入到go/bin目录下。

4. 在当前目录执行protoc-gen-mqant 会自动生成代码

5. 目录解释

- realease.yaml
> 配置目录，插件会根据参数选择合适的文件进行输入和输出（无特殊原因可以不修改）
- proto
> pb文件存放位置
- proto/msg
1. 传输信息存放的目录 **option go_package = "msg/login";** protoc-gen插件会根据创建文件
2. msg目录下是每个服务对应的pb文件。目前插件定死每个服务需要定义一个文件夹。下面是对应的文件

- proto/netmsg/loginsvc

```go
syntax = "proto3";
package loginsvc;
option go_package="netmsg/loginsvc";
import "auto-start/pkg/gen/msg/login";

// 接口
service LoginService {
   rpc  GetUidFromChannel(session gate.Session, req *login.LoginRequest)(rsp *login.LoginResponse,err error)
   rpc  VerifyUid(ctx context.Context,req *login.LoginRequest)(rsp *login.LoginResponse,err error)
}
```

1. package:生成的服务的包名
2. go_package:导出的go文件路径
3. import:导入的文件名，auto-start是你的项目名。pkg/gen/msg/login 是文件目录的名