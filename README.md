# Pod 🌱

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/pod.svg)](https://pkg.go.dev/github.com/brittonhayes/pod)

> Project and client management application for audio professionals

<p align=center>
    <image src="./preview.png" width="500px"/>
</p>

## Installation

Check out the releases for the latest OSX installers and Linux executables.

## Development

### Requirements

- [Taskfile](https://taskfile.dev/#/)
- [Go](https://golang.org/doc/install)
- [NodeJS](https://nodejs.org/en/download/)
- [Wails](https://wails.app/gettingstarted/)

### Download

```shell
# Clone the repositoru
git repo clone brittonhayes/pod

# Navigate to the repository
cd pod

# Download the dependencies
go mod download && cd frontend && npm install && cd ..
```

### Compile from source

```shell
wails build
```
