FROM mhart/alpine-node:10 AS builder
RUN apk add tzdata
ENV TZ=America/Argentina/Buenos_Aires
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

FROM mhart/alpine-node:10
WORKDIR /app
COPY --from=builder /app .
EXPOSE 3000
CMD ["node_modules/.bin/next", "start"]