# goleveldb_http_server



## Introduction

```
Golang leveldb as a KV http server. It supports user-defined tables,
which acts as different namespaces . Can be used in different scene.

As each table act as a database,so I will check periodically whether
a table is used recently, and close unused tables in RefreshAsync().

If you want to deploy several instances on a server,  it is possible
for you to use different ports and different directories to log data.
```


# Example



## Get data

```
curl '127.0.0.1:8880/data?table=2020.03.20&act=get&key=kkk'

After curl, we get the value of key kkk in table 2020.03.20
```

## Del data

```
curl '127.0.0.1:8880/data?table=2020.03.20&act=del&key=kkk'

After curl, we delete the key kkk from the table 2020.03.20
```

## Put data

```
curl --data "123456" '127.0.0.1:8880/data?table=2020.03.20&act=put&key=kkk'

After curl, we insert the key kkk with value 123456 to the table 2020.03.20
```



## Description

```
The above uses date as table. Data is saved in /data/logs/leveldb/db/2020.03.20/, so we can easily use

it to save log and rotate everyday to del old. Table can also be named according to different services. 

One should be care that table should not have ../ or .. to avoid overwrite directory(Never run as root).
```



# Further more work

```
0. 增加读写性能测试。(performance test)

1. 可以增加主从备份,简单地命令传播就可实现。(master-slave)

2. 可以在前面增加一个proxy以实现key分布式。(distributed)
```


Refer : https://github.com/syndtr/goleveldb

Cited : https://github.com/11061055/Transfer (在这里使用过)
