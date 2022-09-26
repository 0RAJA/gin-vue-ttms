FROM node:16 AS builder

WORKDIR /ttms_web/
COPY . .

RUN npm config set registry http://registry.npm.taobao.org
RUN npm install
RUN cp .docker/modules/qrcode.js node_modules/qrcodejs2/qrcode.js
RUN npm run build

FROM nginx:alpine

COPY .docker/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=builder /ttms_web/dist /usr/share/nginx/html