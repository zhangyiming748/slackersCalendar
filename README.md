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
## WebServer(port int)

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

### Upload


```
@ 请求示例
{post} http://127.0.0.1:3306/upload

允许上传2M以下txt文件
```

## 效果
```shell
# 2000次请求 每次并发50个
$ hey -n 2000 -c 50 -m GET http://127.0.0.1:3306/HappyCount

Summary:
  Total:        15.0414 secs
  Slowest:      0.4250 secs
  Fastest:      0.1532 secs
  Average:      0.3758 secs
  Requests/sec: 132.9660

  Total data:   3641870 bytes
  Size/request: 1820 bytes

Response time histogram:
  0.153 [1]     |
  0.180 [11]    |
  0.208 [88]    |■■
  0.235 [50]    |■
  0.262 [0]     |
  0.289 [3]     |
  0.316 [46]    |■
  0.343 [1]     |
  0.371 [9]     |
  0.398 [1424]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.425 [367]   |■■■■■■■■■■


Latency distribution:
  10% in 0.3440 secs
  25% in 0.3871 secs
  50% in 0.3893 secs
  75% in 0.3930 secs
  90% in 0.4054 secs
  95% in 0.4149 secs
  99% in 0.4231 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.1532 secs, 0.4250 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:    0.3756 secs, 0.1532 secs, 0.4249 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200] 2000 responses


```