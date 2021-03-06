#!/bin/bash

# checkout 
cd ~/yource
git pull

# deploy frontend
cd frontend
npm run build
