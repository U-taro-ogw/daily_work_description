FROM node:12.13.0-alpine

WORKDIR /vue_app

COPY package.json ./
COPY yarn.lock ./

RUN apk update && \
    yarn global add @vue/cli && \
    yarn install

COPY . ./

EXPOSE 3000

CMD [ "yarn", "serve" ]