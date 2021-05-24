# Sysbench

Env
* Git commit: 52edb542935827730b1d7fa12ed702427f89feb8
* Macbook 16inch 2.6 GHz 6Cores Intel Core i7 16GB 2667MHz DDR4

Tests 
* mysql: sysbench -> mysql(docker)
* database-mesh: sysbench -> database-mesh -> mysql(docker)

## prepare
```shell
sysbench oltp_common.lua --mysql-host=127.0.0.1 --mysql-port=3306 --mysql-db=test --mysql-user=root --mysql-password=root --table_size=100000 --db-ps-mode=disable --mysql-port=3306 prepare
```

## oltp_read_only
### mysql
```shell
$ for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_read_only --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3306 run | grep 'transactions:' ; done
    transactions:                        2641   (44.01 per sec.)
    transactions:                        4748   (79.10 per sec.)
    transactions:                        8161   (135.95 per sec.)
    transactions:                        12047  (200.70 per sec.)
    transactions:                        14135  (235.32 per sec.)
    transactions:                        16344  (271.85 per sec.)
    transactions:                        16083  (266.23 per sec.)
```

### database-mesh 

```shell
$ for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_read_only --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3308 run | grep 'transactions:' ; done
    transactions:                        2500   (41.66 per sec.)
    transactions:                        3916   (65.25 per sec.)
    transactions:                        5957   (99.22 per sec.)
    transactions:                        7931   (132.07 per sec.)
    transactions:                        10206  (169.90 per sec.)
    transactions:                        12020  (199.85 per sec.)
    transactions:                        11371  (188.30 per sec.)
```


## oltp_read_write
### mysql

``` shell
$ for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_read_write --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3306 run | grep 'transactions:' ; done
    transactions:                        1783   (29.70 per sec.)
    transactions:                        3478   (57.94 per sec.)
    transactions:                        5844   (97.37 per sec.)
    transactions:                        8781   (146.18 per sec.)
    transactions:                        9631   (160.33 per sec.)
    transactions:                        11034  (183.46 per sec.)
    transactions:                        10641  (176.21 per sec.)
    
```

### database-mesh 

``` shell
$ for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_read_write --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3308 run | grep 'transactions:' ; done
    transactions:                        1583   (26.38 per sec.)
    transactions:                        3150   (52.49 per sec.)
    transactions:                        5574   (92.85 per sec.)
    transactions:                        8078   (134.52 per sec.)
    transactions:                        13645  (227.00 per sec.)
    transactions:                        9640   (160.23 per sec.)
    transactions:                        14773  (245.23 per sec.)
```

## oltp_write_only
### mysql

```shell
$ for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_write_only --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3306 run | grep 'transactions:' ; done
    transactions:                        7825   (130.41 per sec.)
    transactions:                        14956  (249.24 per sec.)
    transactions:                        26013  (433.14 per sec.)
    transactions:                        40122  (668.56 per sec.)
    transactions:                        45560  (758.88 per sec.)
    transactions:                        57663  (959.57 per sec.)
```

### database-mesh 

```shell
$ for i in 1 2 4 8 16 32 64 128 ; do sysbench oltp_write_only --num-threads=$i --max-requests=0 --max-time=60 --table-size=100000 --mysql-user=root --mysql-password=root --mysql-host=127.0.0.1 --mysql-db=test --db-ps-mode=disable --mysql-port=3308 run | grep 'transactions:' ; done
    transactions:                        6830   (113.82 per sec.)
    transactions:                        13641  (227.32 per sec.)
    transactions:                        24512  (408.47 per sec.)
    transactions:                        26358  (439.15 per sec.)
    transactions:                        41417  (689.68 per sec.)
    transactions:                        40120  (667.83 per sec.)
    transactions:                        51441  (854.59 per sec.)
```


