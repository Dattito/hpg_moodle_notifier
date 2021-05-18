FROM node:alpine

# For setting right timezone
RUN apk add --no-cache tzdata
WORKDIR /app

ENV NODE_ENV=production

COPY package.json yarn.lock ./

RUN yarn --frozen-lockfile

COPY . .

RUN npm prune --production

ENV NODE_ENV=production POSTGRES_HOST=pghost POSTGRES_PASSWORD=passwd POSTGRES_USER=postgres POSTGRES_DB=postgres \
    MOODLE_HOST=moodlehost MOODLE_TOKEN_SERVICE=moodle_mobile_app SYNC_TABLES=no \
    UPDATE_CRONTAB="" UPDATE_ON_START=0 NO_SCHEDULER=0 TZ=Europe/Berlin

CMD ["node", "./index.js"]