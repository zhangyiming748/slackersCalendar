# slackersCalendar
基于之前创建的库在同一个版本(v1)上实现不同的功能
~~不同版本之间同步功能实在是太让人崩溃了~~
# 函数
### Happy()
#### model.HappyTimer()
#### model.HappyDay()
#### model.Gift()
#### model.AnniversaryDay()
#### model.HappyFinnal()
### Sad()
#### model.SadTimer()
#### model.SadDay()
###WebServer(port int)
# 接口
需要传入端口号开启服务
##WebServer(port int)

### HappyCount
```
@ 请求示例
{get} http://127.0.0.1:3306/HappyCount
@ 返回示例
{
"title": "Happy",
"contents": [
"早上好,摸鱼人!",
"今天是2022年1月11日",
"星期二,还有三天周末",
"上午 10点59分54秒",
],
"ip": "127.0.0.1"
}
```
### SadCount
```
@ 请求示例
{get} http://127.0.0.1:3306/SadCount
@ 返回示例
{
"title": "Happy",
"contents": [
"早上好,开始内卷的一天!",
"今天是2022年1月11日",
"上午11点06分31秒",
"星期二",
"【内卷办】提醒您",
"距离高考还有146天"
],
"ip": "127.0.0.1"
}
```