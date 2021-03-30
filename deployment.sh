#!/bin/bash

# setup env
export PATH=$PATH:/home/ubuntu/.nodenv/shims:/home/ubuntu/.nodenv/bin:/usr/bin/go/bin:

# setup go modules
cd ~/yource/backend
go mod tidy

# フロントエンドはvercelに移管したのでコメントアウト
# # deploy frontend
# cd ../frontend
# npm install 
# npm run build
