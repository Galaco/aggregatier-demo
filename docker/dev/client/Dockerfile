FROM node:8.16.0-alpine

RUN mkdir -p /srv/app/aggregatier
WORKDIR /srv/app/aggregatier

COPY client/package.json /srv/app/aggregatier
COPY client/package-lock.json /srv/app/aggregatier

RUN npm install

COPY client/ /srv/app/aggregatier

CMD [ "npm", "run-script", "serve" ]