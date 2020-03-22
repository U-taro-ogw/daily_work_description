FROM node:12.13.0-alpine

WORKDIR /vue_app

COPY package.json ./
COPY yarn.lock ./

RUN yarn install

COPY . ./

EXPOSE 3000

CMD [ "yarn", "serve" ]