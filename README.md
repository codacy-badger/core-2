[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0a95012e51f941c5abf06f4133b36b77)](https://www.codacy.com/manual/pravinba9495/iborg-core?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=pravinba9495/iborg-core&amp;utm_campaign=Badge_Grade) ![CI](https://github.com/pravinba9495/iborg-core/workflows/CI/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/pravinba9495/iborg-core)](https://goreportcard.com/report/github.com/pravinba9495/iborg-core) [![codecov](https://codecov.io/gh/pravinba9495/iborg-core/branch/master/graph/badge.svg)](https://codecov.io/gh/pravinba9495/iborg-core)

# iborg-core
## Download and Install Go
### Linux, macOS, and FreeBSD tarballs
[Download the archive](https://golang.org/dl/) and extract it into /usr/local, creating a Go tree in `/usr/local/go`. For example:

```shell
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```
Choose the archive file appropriate for your installation. For instance, if you are installing Go version 1.2.1 for 64-bit x86 on Linux, the archive you want is called `go1.2.1.linux-amd64.tar.gz`.

(Typically these commands must be run as root or through sudo.)

Add `/usr/local/go/bin` to the PATH environment variable. You can do this by adding this line to your `/etc/profile` (for a system-wide installation) or `$HOME/.profile`:

```shell
export PATH=$PATH:/usr/local/go/bin
```
**Note**: changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source `$HOME/.profile`.

### macOS package installer

[Download the package file](https://golang.org/dl/), open it, and follow the prompts to install the Go tools. The package installs the Go distribution to `/usr/local/go`.

The package should put the `/usr/local/go/bin directory` in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

### Windows

Open the [MSI file](https://golang.org/dl/) and follow the prompts to install the Go tools. By default, the installer puts the Go distribution in `C:\Go`.

The installer should put the `C:\Go\bin` directory in your PATH environment variable. You may need to restart any open command prompts for the change to take effect.

### Zip archive

[Download the zip file](https://golang.org/dl/) and extract it into the directory of your choice (we suggest `C:\Go`).

Add the `bin` subdirectory of your Go root (for example, `C:\Go\bin`) to your PATH environment variable.

### Setting environment variables under Windows

Under Windows, you may set environment variables through the "Environment Variables" button on the "Advanced" tab of the "System" control panel. Some versions of Windows provide this control panel through the "Advanced System Settings" option inside the "System" control panel.

## Setting up the repository
### Clone the repository

Use `git` to clone the repository:

```shell
git clone https://github.com/iborg-ai/iborg-core.git
```

### Install the dependencies

The package uses Go Modules to manage dependencies. Install dependencies by typing:

```shell
go mod download
```

### Run the package

Run the package by typing:

```shell
go run ./src/main
```

### Run tests

```shell
go test -v ./src/* -race -coverprofile=coverage.txt -covermode=atomic
```

### Build the package

Generate a binary build for your operating system by typing:

```shell
# For Windows
env GOOS=windows GOARCH=amd64 go build -o ./bin/windows/core ./src/main

# For macOS
env GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/core ./src/main

# For Linux
env GOOS=linux GOARCH=amd64 go build -o ./bin/linux/core ./src/main
````

## License
MIT licensed. See the LICENSE file for details.
