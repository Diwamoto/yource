cd docker
docker-compose down && docker-compose up -d --build
cd ../frontend
yarn dev