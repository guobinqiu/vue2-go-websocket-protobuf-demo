# backend

### 安装组件

> https://github.com/guobinqiu/grpc-exercise/blob/master/go-as-serverside/readme.md

### 根据 chat.proto 生成代码

```
protoc --go_out=. --go_opt=paths=source_relative chat/chat.proto
```

| 参数                             | 含义                                                                                      |
| -------------------------------- | ----------------------------------------------------------------------------------------- |
| `protoc`                         | Protobuf 编译器命令，用于将 `.proto` 文件编译为目标语言（如 Go）的源代码。                |
| `--go_out=.`                     | 指定生成的 Go 源文件输出到当前目录，并调用 `protoc-gen-go` 插件生成对应的 `.pb.go` 文件。 |
| `--go_opt=paths=source_relative` | 设置输出路径为与 `.proto` 文件相对的路径，而不是根据 `go_package` 中定义的路径生成结构。  |
| `chat/chat.proto`                | 需要被编译的 `.proto` 文件路径。                                                          |


```
protoc --go_out=. chat/chat.proto
```
不加`--go_opt=paths=source_relative` 会根据 `option go_package = "github.com/guobinqiu/vue2-go-websocket-protobuf-demo/chat";` 产生对应目录, 不推荐
