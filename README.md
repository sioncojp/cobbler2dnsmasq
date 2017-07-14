# cobbler2dnsmasq

## Overview
* cobblerのprofileから下記を読み取って`/etc/cobbler/dnsmasq.template`を作ります
  * 一番最後のMACアドレス
  * 正規表現に引っかかった+ignore_IPを除外したIPアドレス
  * 正規表現に引っかかったホスト名

## Test Run

```shell
$ git clone github.com/sioncojp/cobbler2dnsmasq
$ make deps
$ make run

```

* `tmp/dnsmasq.template` が生成されます。ファイルの中身は下記。

```
# sample
dhcp-host=AA:AA:AA:AA:AA:33,fuga,192.168.0.2,infinite
dhcp-host=AA:AA:AA:AA:AA:11,hoge,192.168.0.1,infinite
```
