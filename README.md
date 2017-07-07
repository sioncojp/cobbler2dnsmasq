# cobbler2dnsmasq

## Overview
* cobblerのprofileから`/etc/cobbler/dnsmasq.template`を作ります
* MACアドレスは一番最後に書かれてるものを取ってきます

## Test Run

```shell
$ git clone github.com/sioncojp/cobbler2dnsmasq
$ dep ensure
$ make run

```

* `tmp/dnsmasq.template` が生成されます。ファイルの中身は下記。

```
# sample
dhcp-host=AA:AA:AA:AA:AA:33,fuga,192.168.0.2,infinite
dhcp-host=AA:AA:AA:AA:AA:11,hoge,192.168.0.1,infinite
```
