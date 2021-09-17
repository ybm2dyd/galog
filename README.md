# galog
```
             _             
  __ _  __ _| | ___   __ _ 
 / _` |/ _` | |/ _ \ / _` |
| (_| | (_| | | (_) | (_| |
 \__, |\__,_|_|\___/ \__, |
 |___/               |___/ 

```

logger sdk for Golang

## 使用方法
参考 example/demo.go

1. 初始化sdk `err := galog.Init(logpath, rollingtime, rollinginterval)`。其中，logpath 指定写日志路径，rollingtime 指定日志切割的时间单位，rollinginterval 指定日志切割的间隔
2. 新建对象 `playerLogout := new(galog.Playerlogout)`
3. 设置字段 `playerLogout.EventTime = "2020-10-19 20:23:13"`
4. 输出日志 `err = playerLogout.Log()`
5. 退出时关闭日志文件 `galog.Clean()`