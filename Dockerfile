########-- Build stage --########

FROM golang:1.19.3-alpine as builder-prod

RUN apk update
RUN apk add git

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -v -o app .

########-- Prod image --########

FROM alpine:3.16 as prod
WORKDIR /opt
COPY --from=builder-prod /opt/app .
CMD [ "/opt/app" ]