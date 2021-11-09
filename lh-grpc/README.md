1. 安装命令行工具 Protocol Buffers
   ```
   brew install protobuf
   
   或者手动安装
   
   wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-all-3.6.1.tar.gz
   tar zxvf protobuf-all-3.6.1.tar.gz && cd protobuf-3.6.1/
   ./configure && make && make install
   ```


2. 安装GO依赖包
   ```
   go get google.golang.org/grpc
   go get github.com/golang/protobuf/protoc-gen-go
   ```

```protobuf
syntax = "proto3"; // 指定proto版本

package proto;     // 指定包名

/**
指定golang包名
option go_package = "path;name";
path 表示生成的go文件的存放地址，会自动生成目录的。
name 表示生成的go文件所属的包名

必须这么写,否则报错: --go_out: protoc-gen-go: Plugin failed with status code 1.
 */
option go_package = "./;simple";

// 定义Hello服务
service Hello {
  // 定义SayHello方法
  rpc SayHello(HelloRequest) returns (HelloReply) {}
}

// HelloRequest 请求结构
message HelloRequest {
  string name = 1;
}

// HelloReply 响应结构
message HelloReply {
  string message = 1;
}
```

编译proto文件

```
protoc -I . --go_out=plugins=grpc:. ./hello.proto

在 proto 文件夹下执行如下命令：

protoc --go_out=plugins=grpc:. *.proto
plugins=plugin1+plugin2：指定要加载的子插件列表
我们定义的 proto 文件是涉及了 RPC 服务的，而默认是不会生成 RPC 代码的，因此需要给出 plugins 参数传递给 protoc-gen-go，告诉它，请支持 RPC（这里指定了 gRPC）

–go_out=.：设置 Go 代码输出的目录
该指令会加载 protoc-gen-go 插件达到生成 Go 代码的目的，生成的文件以 .pb.go 为文件后缀

: （冒号）
冒号充当分隔符的作用，后跟所需要的参数集。如果这处不涉及 RPC，命令可简化为：

protoc --go_out=. *.proto
注：建议你看看两条命令生成的 .pb.go 文件，分别有什么区别
```


如果出现下面的报错，是因为 go 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。 如果想兼容之前的方式，需要设置环境变量 GODEBUG 为 x509ignoreCN=0。
```
time="2020-11-18T12:54:48+08:00" level=fatal msg="rpc error: code = Unavailable desc = connection error: desc = \"transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0\""
```

```
export GODEBUG="x509ignoreCN=0"
```

