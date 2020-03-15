FROM golang:1.14.0-alpine3.11 as builder 
COPY main.go .
RUN apk add git
RUN go get -u github.com/gin-gonic/gin
RUN go build -o /app main.go

FROM alpine:latest
EXPOSE 8085 
CMD ["./app"]
COPY --from=builder /app .