# cobbler2dnsmasq

## Overview
* Create `/etc/cobbler/dnsmasq.template` from cobbler profiles
* After creating, run command `cobbler sync`

## Test Run

```shell
$ git clone github.com/sioncojp/cobbler2dnsmasq
$ cp -rp tmp/`cobbler profiles directory`/
$ make run
```
