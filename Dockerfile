# build
FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o ip2region
RUN apk update && apk add tzdata

# deploy
FROM alpine
ARG VERSION
ENV VERSION $VERSION
WORKDIR /app
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
COPY --from=build /app/ip2region /app/ip2region
EXPOSE 8080
CMD [ "/app/ip2region" ]
