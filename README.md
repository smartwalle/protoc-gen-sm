## 功能

### 为 oneof 类型字段添加 New 函数。

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

### 为 oneof 类型字段添加相关方法

#### Number() int32

获取该字段的编号

#### String() string

获取该字段的名称

#### Filename() string

获取该字段类型声明所在的 proto 文件名

## 安装

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/smartwalle/protoc-gen-sm@latest
```

或者将本项目下载到本地之后，进入项目所在目录执行 **go install** 命令。

## 使用

```shell
protoc  --go_out=. --sm_out=. --go-grpc_out=. ./*.proto
```

## 参考

[https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code)