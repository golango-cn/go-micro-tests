#### 下载micro

```go
go get -u -v github.com/go-log/log
go get -u -v github.com/gorilla/handlers
go get -u -v github.com/gorilla/mux
go get -u -v github.com/gorilla/websocket
go get -u -v github.com/mitchellh/hashstructure
go get -u -v github.com/nlopes/slack
go get -u -v github.com/pborman/uuid
go get -u -v github.com/pkg/errors
go get -u -v github.com/serenize/snaker
go get -u -v github.com/hashicorp/consul
go get -u -v github.com/miekg/dns
go get -u -v github.com/micro/micro
go get -u -v github.com/micro/go-micro
```

#### 编译安装
```go
cd $GOPATH/src/github.com/micro/micro
go build main.go
```

#### 插件安装
```go
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
```

#### 创建服务

```

micro new --type "srv" cc-go-micro/ccserver
Creating service go.micro.srv.ccserver in cc-go-micro/ccserver

.
├── .
│   ├── main.go
│   ├── generate.go
│   ├── plugin.go
│   ├── Dockerfile
│   ├── Makefile
│   ├── README.md
│   ├── .gitignore
│   └── go.mod
├── handler
│   └── ccserver.go
├── subscriber
│   └── ccserver.go
└── proto\ccserver
    └── ccserver.proto


download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro/v2

compile the proto file ccserver.proto:

cd cc-go-micro/ccserver
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/ccserver/ccserver.proto

```

```
λ micro new --type "api" cc-go-micro/ccclient
Creating service go.micro.api.ccclient in cc-go-micro/ccclient

.
├── .
│   ├── main.go
│   ├── generate.go
│   ├── plugin.go
│   ├── Makefile
│   ├── Dockerfile
│   ├── README.md
│   ├── .gitignore
│   └── go.mod
├── client
│   └── ccclient.go
├── handler
│   └── ccclient.go
└── proto\ccclient
    └── ccclient.proto


download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro/v2

compile the proto file ccclient.proto:

cd cc-go-micro/ccclient
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/ccclient/ccclient.proto

```

###　修改文件

- 拷贝ccserver下的proto文件到ccclient下
- 修改ccclient/handler.go

```go
// call the backend service
ccclientClient := ccserver.NewCcserverService("go.micro.srv.ccserver", client.DefaultClient)
rsp, err := ccclientClient.Call(context.TODO(), &ccserver.Request{
    Name: request["name"].(string),
})
if err != nil {
    http.Error(w, err.Error(), 500)
    return
}
```
　

#### 运行服务端
```
go run main.go
2020-03-02 19:51:34 Starting [service] go.micro.srv.ccserver
2020-03-02 19:51:34 Server [grpc] Listening on [::]:59050
2020-03-02 19:51:34 Broker [eats] Connected to [::]:59052
2020-03-02 19:51:34 Registry [mdns] Registering node: go.micro.srv.ccserver-9a5f1512-8ca0-453b-80c2-478e0a0ae7c2
2020-03-02 19:51:35 Subscribing to topic: go.micro.srv.ccserver

```

#### 运行客户端
```
go run main.go
2020-03-02 19:51:40 Listening on [::]:59057
```

#### 打开浏览器，输入上面端口号的本地地址

```go
http://localhost:59057
```