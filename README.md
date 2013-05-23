**本程序 forked from https://github.com/msbranco/goconfig**

###为了满足个人需求，做了简化和简单修改

1. 去除了原版的写入配置文件的功能，只需要读取配置文件
2. 去除了原版的配置内容可相互引用的处理，不需要那么复杂，KEY->VALUE就可以了
3. 修改了firstIndex函数的处理，把二次循环改为 strings.Index() 返回pos
4. 值获取时自动去除两端的引号
5. 重写了测试用例，简化为四个测试：读取配置文件，获取字符串值，获取整数值，获取浮点数，获取布尔值

###简单重复下程序的特性：

1. 支持linux下最为传统的配置文件写法
2. 支持多个section片段
3. 值配置可以使用等号和冒号“=:”
4. 布尔值支持多种格式，包括：y/n,yes/no,true/false,no/off
5. 可选的值引号包裹

###程序中已经提交了测试用的配置文件示例，如：

    [redis]
    redisAddr = "192.168.1.80:6379"
    redisDb = 0
    redisList = "ltest"
    
    [log]
    logOpen = no
    logFile = "/var/log/test.log"
    logDays = 14
    logSize = 1.5
