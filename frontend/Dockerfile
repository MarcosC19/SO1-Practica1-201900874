FROM node

WORKDIR /app

COPY package*.json ./

COPY . .

RUN npm i

ENV REACT_APP_HOSTIP=192.168.1.11

CMD ["npm", "run", "start"]