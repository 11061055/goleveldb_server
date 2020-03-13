# goleveldb_server

Golang leveldb as a KV server. It Supports user-defined  file,

which acts as a namespace . Cat be used in different scene.

## Put data

curl --data "123456" '127.0.0.1:8880/data?act=put&key=kkk&file=2020.03.20'


After curl, file 2020.03.20 contains a key named kkk and with value 123456

## Get data

curl '127.0.0.1:8880/data?act=get&key=kkk&file=2020.03.20'


After curl, we get the value of key kkk in file 2020.03.20

## Del data

curl '127.0.0.1:8880/data?act=del&key=kkk&file=2020.03.20'


After curl, the key kkk is deleted  from  file 2020.03.20

## Description

the above scene name file as date, data will be saved in /data/logs/leveldb/db/2020.03.20/ and we can easily use

it to save log file and rotate everyday to delete old one, file can also be distinguished according to different 

service. one should be care that file could not contain ../ or .. to avoid overwrite directory(Never run as root).

## Further more work

0. 增加读写性能测试。(performance test)
1. 可以增加主从备份,简单地命令传播就可实现。(master-slave)
2. 可以在前面增加一个proxy以实现key分布式。(distributed)
