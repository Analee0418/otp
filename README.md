# README #

 一键

### 目的 ###
运行命令自动将 OTP 验证码复制到剪切板，不用再看手机
当然，如果是资金安全相关还是会用手机的。。

* hard code secret
* 所以需要你自己将你的 secret 添加到源码中，并编译它，然后不要在任何地方暴露
* 当然第三条有些门槛
* 执行之后可以自动拿到你要的 验证码，并自动放入剪切板
* 配置 Spotlight 或 Alfred 简单快捷！！！
* 减少拿起手机的次数，工作中拿起手机会中断你的工作，而且！！！！ 很有可能你就会在拿起手机后点开 某音、某博、某推。
* 要觉得有用就 clone 下来吧 ！！

### 怎么用 ###

* clone 代码，如果没有 golang 环境，，那你就去装一个！！
* 修改 main.py 哦对不起，是 main.go 里的 [line:21](https://github.com/Analee0418/otp/blob/05c5668863b491d051bc4ce3288c5da97b8e0d78/main.go#L21) 加 别名 和 对应的 secret
* 然后 build
* 然后 ./opt $别名，然后 command + v 试试
