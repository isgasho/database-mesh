#! /bin/bash

conprof all --config.file configs/config.yaml  --http-address :10902 --storage.tsdb.path ./data
