#! /bin/bash

sysbench oltp_common.lua --mysql-host=127.0.0.1 --mysql-port=3306 --mysql-db=test --mysql-user=root --mysql-password=root --table_size=100000 --db-ps-mode=disable --mysql-port=3306 prepare

