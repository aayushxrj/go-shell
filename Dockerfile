FROM golang:1.22.2

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o myshell ./cmd/myshell/main.go 

EXPOSE 8080

CMD [ "./myshell" ]

# docker run -it aayushxrj/go-shell:latest