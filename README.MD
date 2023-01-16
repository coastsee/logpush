## Log push
🤔️还在用命令行查日志？不如试试这个！它能帮你把日志根据长度或者时间分批次推送到搜索引擎上去，并在web上检索它们，这一切非常简单。

️🤔️Still checking logs using the command line? Why don't you try this! It makes it easy to push your logs to search engines in batches, depending on length or time, and retrieve them on the Web.

🐮我们已经支持了zinc搜索引擎，这意味着你可以开箱即用，当然，你也完全可以根据下面的接口规范编写一个自己常用的搜索引擎实现，例如：es，如果你想知道怎么做，可以查看engines/zinc.go，非常简单。

🐮We already support the Zinc search engine, which means you can use it right out of the box. Of course, you can also write your own common search engine implementation according to the following interface specification, such as es. If you want to know how to do this, you can check out Engines/zinc.go file.

```
type Engine interface {
    Flush([]string) error
}
```

## 快速开始 Quick start

```shell
go get github.com/coastsee/logpush
```

```
var logPush = LogPush{
    MaxPoolLength: 5,
    PushDuration:  time.Minute,
    Engine: engines.Zinc{
      Index:"test",
      Url: "http://localhost",
      Username: "your username",
      Password: "your password",
    },
}

err := logPush.Push("log content")
if err != nil {
    fmt.Println(err)
}

// for zapcore
core := zapcore.NewCore(
    encoder,
    zapcore.NewMultiWriteSyncer(zapcore.AddSync(&toZine{Level: logLevel.String()}),......),
    zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == logLevel
    }), // 日志级别
)
```
## 参数说明 Parameter

#### @maxPoolLength
日志池最大长度，超过就推送到引擎

The maximum length of the log pool is pushed to the engine

#### @pushDuration
两次推送间隔时间，超过就推送到引擎

The interval between two pushes will be pushed to the engine

#### @engine
你可以创建一个自己的engine，符合logpush.Engine interface即可。

You can create your own engine that conforms to logpush.Engine interface.

## Supported search engines
没有限制，我们已经实现了zinc引擎的推送。

No limit, We have implemented zinc engine.
