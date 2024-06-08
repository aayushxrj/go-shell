# Go Shell

Go Shell is a [POSIX](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html) compliant shell implemented in Go. It's capable of interpreting shell commands, running external programs, and executing built-in commands. 

## Built-in commands

- [x] echo
- [x] type
- [x] exit
- [X] pwd
- [X] cd

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Running on local machine 

#### Prerequisites

- Go 1.22.2 or later

Clone the repository:

```sh
git clone https://github.com/aayushxrj/go-shell
cd go-shell
```

Build the application:

```sh
go build -o myshell ./cmd/myshell/main.go 
```

Run the application:

```sh
./myshell
```

### Running with Docker

You can also run Go Shell inside a Docker container.

Directly pull and run from Docker Hub (recommended):

```sh
docker-compose run myshell
```

or Build and run the Docker image locally:

```sh
docker build -t aayushxrj/go-shell:latest .
```

```sh
docker run -it aayushxrj/go-shell:latest
```

## License

This project is licensed under the MIT License.