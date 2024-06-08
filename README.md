# Go Shell

Go Shell is a [POSIX](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html) compliant shell implemented in Go. It's capable of interpreting shell commands, running external programs, and executing built-in commands. 

## Built-in commands

- [x] echo - Display a line of text.
- [x] cd - Change the current directory.
- [x] pwd - Print the current directory.
- [x] type - Display the type of command (built-in, alias, file, etc.).
- [x] exit - Exit the shell.

- [ ] help - Display help for built-in commands.
- [ ] | - Piping operator to pass the output of one command as input to another.
- [ ] >, >> - Redirection operators to direct output to a file (overwrite or append).
- [ ] history - Show command history.

- [ ] ls - List directory contents.
- [ ] cp - Copy files or directories.
- [ ] mv - Move/rename files or directories.
- [ ] rm - Remove files or directories.
- [ ] mkdir - Create directories.
- [ ] rmdir - Remove empty directories.
- [ ] touch - Change file timestamps or create empty files.

- [ ] cat - Concatenate and display file content.
- [ ] less - View file content page by page.
- [ ] head - Display the beginning of a file.
- [ ] tail - Display the end of a file.
- [ ] nano or vim - Text editors (if you want to integrate an editor).

- [ ] ps - Display current processes.
- [ ] kill - Terminate a process by ID.
- [ ] top - Display system resource usage (if you want to integrate a basic version).

- [ ] ping - Check network connectivity.
- [ ] curl or wget - Download files from the web (basic version).

- [ ] df - Display disk space usage.
- [ ] du - Estimate file space usage.
- [ ] free - Display memory usage.

- [ ] env - Display environment variables.
- [ ] export - Set environment variables.
- [ ] whoami - Display the current user.
- [ ] su - Switch user (requires proper permission handling).

- [ ] echo - Display a line of text.
- [ ] alias - Create an alias for a command.
- [ ] source - Execute commands from a file in the current shell.

- [ ] chmod - Change file permissions.
- [ ] chown - Change file owner and group.
- [ ] ln - Create hard and symbolic links.
- [ ] find - Search for files in a directory hierarchy.
- [ ] grep - Search text using patterns.

- [ ] man - Display the manual for commands.
- [ ] date - Display or set the system date and time.
- [ ] hostname - Show or set the system's hostname.
- [ ] Enable autocompletion for faster and more efficient command execution.

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