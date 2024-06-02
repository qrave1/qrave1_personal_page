FROM huecker.io/library/golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]
