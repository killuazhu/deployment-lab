 FROM golang:1.16.0-alpine
 COPY main.go .
 RUN go build -o web main.go
 EXPOSE 8090
 CMD ["./web"]