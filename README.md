# goleveldb_http_server


## Description

Golang leveldb as a KV server. It Supports user-defined table,

which acts as a namespace . Cat be used in different scene.

## Get data

curl '127.0.0.1:8880/data?table=2020.03.20&act=get&key=kkk'


After curl, we get the value of key kkk in table 2020.03.20

## Del data

curl '127.0.0.1:8880/data?table=2020.03.20&act=del&key=kkk'


After curl, we delete the key kkk from the table 2020.03.20

## Put data

curl --data "123456" '127.0.0.1:8880/data?table=2020.03.20&act=put&key=kkk'


After curl, we insert the key kkk with value 123456 to the table 2020.03.20


## Description

The above scene use date as table. Data will be saved in /data/logs/leveldb/db/2020.03.20/, so we can easily use

it to save log file and rotate everyday to delete old one. Table can also be distinguished according to different 

service. One should be care that table should not have ../ or .. to avoid overwrite directory (Never run as root).

## Further more work

0. 增加读写性能测试。(performance test)
1. 可以增加主从备份,简单地命令传播就可实现。(master-slave)
2. 可以在前面增加一个proxy以实现key分布式。(distributed)
