FROM node:16

WORKDIR /usr/app

COPY package.json .

RUN npm install

COPY . .

EXPOSE 3003

ENTRYPOINT [ "node", "src/server.js" ]
