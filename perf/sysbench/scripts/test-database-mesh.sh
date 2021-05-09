#! /bin/bash

for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_read_write --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3306 run | grep 'transactions:' ; done
