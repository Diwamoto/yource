#!/bin/bash

# setup env
source ~/.profile

# checkout 
cd ~/yource
git pull

# setup go modules
cd backend
go mod tidy

# deploy frontend
cd ../frontend
npm install 
npm run build
