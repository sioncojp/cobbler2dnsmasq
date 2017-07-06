# cobbler2dnsmasq

## Overview
* cobblerのprofileから`/etc/cobbler/dnsmasq.template`を作ります
* MACアドレスは一番最後に書かれてるものを取ってきます

## Test Run

```shell
$ git clone github.com/sioncojp/cobbler2dnsmasq
$ dep ensure
$ make run

# check tmp/dnsmasq.template
```
