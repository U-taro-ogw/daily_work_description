FROM node:12.13.0-alpine

WORKDIR /react_app

COPY package*.json ./

RUN npm install && yarn install

COPY . ./

CMD [ "yarn", "start" ]