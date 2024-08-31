FROM golang:1.22.1-alpine3.19
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
CMD ["echo $ENV_CONTENT > /etc/setting.yaml && ./main"]
