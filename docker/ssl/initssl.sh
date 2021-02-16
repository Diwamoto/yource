brew install mkcert
mkcert -install
mkdir ./keys
mkcert -cert-file ./keys/server.crt -key-file ./keys/server.key $(cat ./hostnames)