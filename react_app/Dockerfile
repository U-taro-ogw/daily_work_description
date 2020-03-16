FROM node:12.13.0-alpine

WORKDIR /react_app/daily_work_description

COPY daily_work_description/package*.json ./

RUN npm install && yarn install

COPY . ./

CMD [ "yarn", "start" ]