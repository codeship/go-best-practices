FROM node:latest

ADD . /repo

WORKDIR /repo

RUN npm install

ENTRYPOINT ["npm", "test"]
