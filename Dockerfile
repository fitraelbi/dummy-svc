FROM golang:alpine as build-env

ARG SERVICE_NAME=dummy-svc

RUN mkdir /service
ADD . /service/

WORKDIR /service
RUN apk add --no-cache git
RUN go build  -o ${SERVICE_NAME} ./app/

FROM alpine
ARG SERVICE_NAME=dummy-svc
WORKDIR /service
COPY --from=build-env /service/${SERVICE_NAME}    /service/${SERVICE_NAME}

EXPOSE 7777

ENTRYPOINT ["/service/dummy-svc"]