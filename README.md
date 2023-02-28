# uestc_keepnet
uestc电信保持连接

build.bat 交叉编译armv7架构可执行文件

run.bat 在Windows上运行时，把go的环境变量设置回来

## 使用方式

抓包或者f12获取流量包，写入配置文件，三个需要配置的，注意url解码一次，在线解码就行，例如：
http://www.jsons.cn/urlencode/

### 原理
每隔一段时间保持一下连接，不在意的话可以一分钟一次，对电脑性能没啥影响

### windows ： 

此电脑（右键）-> 管理 -> 任务计划程序

添加一个计划任务，每小时执行一下（自行选择）

PATH_TO_KEEPLOGIN/keepLogin.exe

### Linux ： 

crontab 

0 * * * * PATH_TO_KEEPLOGIN/keepLogin
