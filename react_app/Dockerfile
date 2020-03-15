FROM node:12.13.0-alpine

WORKDIR /react_app
# 1. workdirではなく、mkdirでディレクトリのみ作成
# 2. daily_work_description/package*json をCOPY
# 3. yarn install ? npm install ?を実行
# 4. 全体をCOPY
# *どっかでWORKDIRの指定が必要かも
COPY package*.json ./

RUN npm install && yarn install

COPY . ./

CMD [ "yarn", "start" ]