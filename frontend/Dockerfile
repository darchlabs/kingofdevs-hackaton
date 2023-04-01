FROM node:16

WORKDIR /usr/src/app

COPY ./frontend/package.json package.json
COPY ./frontend/package-lock.json package-lock.json

RUN npm ci

COPY ./frontend .

RUN npm run build

ENV NODE_ENV=production

CMD ["npm", "run", "start"]


