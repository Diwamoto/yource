#!/bin/bash

# setup env
export PATH=$PATH:/home/ubuntu/.nodenv/shims:/home/ubuntu/.nodenv/bin

# setup go modules
cd ~/yource/backend
go mod tidy

# deploy frontend
cd ../frontend
npm install 
npm run build
