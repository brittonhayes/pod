# Pod 🌱

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/pod.svg)](https://pkg.go.dev/github.com/brittonhayes/pod)
[![Test](https://github.com/brittonhayes/pod/actions/workflows/test.yml/badge.svg)](https://github.com/brittonhayes/pod/actions/workflows/test.yml)

> Productivity application for audio professionals

<p align=center>
    <image src="./appicon--512.png" width="312px"/>
</p>

## Features ✨

- Simple database of all clients, projects, and sound bytes
- Speedy search for all your client documents
- Hub to track the status of your recording sessions
- Batch audio processing (normalize, rename, and more)
- Keep your computer's projects/clients folders clean and organized

## Installation 📥

Available for Mac and Linux.

Check out the [releases](https://github.com/brittonhayes/pod/releases) for the latest installer.

## Preview

<p align=center>
    <image src="./preview.png" width="500px"/>
</p>

## Development 🛠️

### Requirements

- [Taskfile](https://taskfile.dev/#/)
- [Go](https://golang.org/doc/install)
- [NodeJS](https://nodejs.org/en/download/)
- [Wails](https://wails.app/gettingstarted/)

### Download

```shell
# Clone the repositoru
gh repo clone brittonhayes/pod

# Navigate to the repository
cd pod

# Download the dependencies
go mod download && cd frontend && npm install && cd ..
```

### Compile pod from source

```shell
wails build
```
