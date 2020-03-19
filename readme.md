# wechat-sender

Nightingale的理念，是将告警事件扔到redis里就不管了，接下来由各种sender来读取redis里的事件并发送，毕竟发送报警的方式太多了，适配起来比较费劲，希望社区同仁能够共建。

这里提供一个微信的sender，参考了[https://github.com/yanjunhui/chat](https://github.com/yanjunhui/chat)，具体如何获取企业微信信息，也可以参看yanjunhui这个repo

*因为个人没有企业微信服务号，没法测试，这个代码未经过测试，希望有账号的朋友可以帮忙测试一下并将结果反馈给我，我的微信是 cnperl *

## compile

```bash
cd $GOPATH/src
mkdir -p github.com/n9e
cd github.com/n9e
git clone https://github.com/n9e/wechat-sender.git
cd wechat-sender
go build
```

如上编译完就可以拿到二进制了。

## configuration

直接修改etc/wechat-sender.yml即可

## pack

编译完成之后可以打个包扔到线上去跑，将二进制和配置文件打包即可：

```bash
tar zcvf wechat-sender.tar.gz wechat-sender etc/wechat-sender.yml etc/message.tpl
```

## test

配置etc/wechat-sender.yml，相关配置修改好，我们先来测试一下是否好使， `./wechat-sender -t <toUser>`，程序会自动读取etc目录下的配置文件，发一个测试消息给`toUser`

## run

如果测试发送没问题，扔到线上跑吧，使用systemd或者supervisor之类的托管起来，systemd的配置实例：


```
$ cat wechat-sender.service
[Unit]
Description=Nightingale wechat sender
After=network-online.target
Wants=network-online.target

[Service]
User=root
Group=root

Type=simple
ExecStart=/home/n9e/wechat-sender
WorkingDirectory=/home/n9e

Restart=always
RestartSec=1
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
```