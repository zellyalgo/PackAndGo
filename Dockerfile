FROM golang:1.17-alpine

ENV CGO_ENABLED 0

COPY go.mod go.sum /app/
COPY cmd /app/cmd/
COPY pkg /app/pkg/
COPY resources /app/resources

WORKDIR /app

RUN go build -o trip cmd/main.go

EXPOSE 8090

CMD [ "/app/trip" ]