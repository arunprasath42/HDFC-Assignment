FROM golang:1.16-alpine

WORKDIR /go/src/github.com/arunprasath42/HDFC-Assignment/api
COPY . .
RUN go mod download && go mod verify
# update the dependencies
RUN go mod tidy
RUN go run main.go
EXPOSE 8080

CMD ["go", "run", "main.go"]
