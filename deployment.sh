#!/bin/bash

# setup env
source ~/.profile

# setup go modules
cd ~/yource/backend
go mod tidy

# deploy frontend
cd ../frontend
npm install 
npm run build
