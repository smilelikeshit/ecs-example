FROM golang:1.14-alpine as builder 
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .



FROM alpine:latest 

ENV PORT=${PORT}
WORKDIR /app
COPY --from=builder /app/main /app
EXPOSE $PORT
ENTRYPOINT [ "./main" ]
