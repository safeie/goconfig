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
5. 支持使用井号和分号“#;”做注释
6. 可选的值引号包裹
7. 自动忽略空行

###程序中已经提交了测试用的配置文件示例，如：

    ;some comments
    [redis]
    redisAddr = "192.168.1.80:6379"
    redisDb = 0
    redisList = "ltest"
    
    [log]
    logOpen = no
    logFile = "/var/log/test.log"
    logDays = 14
    logSize = 1.5

###使用非常简单：

    package main

	import (
		"fmt"
		"github.com/9466/goconfig"
	)

	func main() {
		c, err := goconfig.ReadConfigFile("t.conf")
		if err != nil {
			fmt.Println(err.Error())
		}
		//  fmt.Println(c)
		sv, err := c.GetString("redis", "redisAddr")
		fmt.Println(sv)
	}
    
**读取字符串：**

    c.GetString(section string, option string) string
    
**读取整数值：**

    c.GetInt64(section string, option string) int64
    
**读取浮点数：**

    c.GetFloat(section string, option string) float64
    
**读取布尔值：**

    c.GetBool(section string, option string) bool
    
###注意：

如果配置字段不包括在任何section中，也就是没写section，程序会自动写入 ***[default]*** 片段。
