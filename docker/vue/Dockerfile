FROM node:10.11-alpine
WORKDIR /app
RUN set -eux && \
  apk update && \
  apk add --no-cache git curl bash vim && \
  npm install && \
  npm install -g @vue/cli && \
  npm install axios --save
