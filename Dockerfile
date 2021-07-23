# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && go build  -o http-responder main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/http-responder /app/
ENTRYPOINT ./http-responder
