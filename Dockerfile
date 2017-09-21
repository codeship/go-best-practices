FROM node:latest

ADD . /repo

WORKDIR /repo

RUN npm install

CMD ["npm", "test"]
