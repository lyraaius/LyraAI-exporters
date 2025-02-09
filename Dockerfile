FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN sh ./build.sh

RUN mv ./output /lyraAiX-exporters

FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache bash && apk add --no-cache tzdata

COPY --from=builder /lyraAiX-exporters /lyraAiX-exporters

ENV GO_ENV=prod

WORKDIR /lyraAiX-exporters

CMD ["bash", "bootstrap.sh"]
