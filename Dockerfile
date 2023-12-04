FROM golang:alpine as build-env

ARG SERVICE_NAME=fiber_name

RUN mkdir /app
ADD . /app/

WORKDIR /app
RUN apk add --no-cache git
RUN go build  -o ${SERVICE_NAME} .

FROM alpine
ARG SERVICE_NAME=fiber_name
WORKDIR /app
COPY --from=build-env /app/${SERVICE_NAME}    /app/${SERVICE_NAME}

EXPOSE 5000

ENTRYPOINT ["/app/fiber_name"]