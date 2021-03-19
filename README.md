## 功能

### 用于生成 oneof 类型字段的 New 函数。

```proto
message Message {
    oneof Body {
        Heartbeat Heartbeat = 1;
        UserInfo UserInfo = 2;
    }
}

message Heartbeat {
}

message UserInfo {
}
```

如上 proto 文件，会额外生成以下两个方法：

```go
func NewMessageHeartbeat() *Message_Heartbeat {
	var m = &Message_Heartbeat{}
	m.Heartbeat = &Heartbeat{}
	return m
}

func NewMessageUserInfo() *Message_UserInfo {
	var m = &Message_UserInfo{}
	m.UserInfo = &UserInfo{}
	return m
}
```

## 安装

```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/smartwalle/protoc-gen-sm
```

## 使用

```shell
protoc  --go_out=. --sm_out=. --go-grpc_out=. ./*.proto
```

## 参考

[https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code)