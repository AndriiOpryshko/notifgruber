FROM golang:1.11.1-alpine3.8 as build-env
RUN mkdir /notifgruber
WORKDIR /notifgruber
COPY go.mod .
COPY go.sum .


RUN apk add --update --no-cache ca-certificates git
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/notifgruber

FROM scratch
COPY --from=build-env /go/bin/notifgruber /go/bin/notifgruber
COPY config.yml .
ENTRYPOINT ["/go/bin/notifgruber"]
