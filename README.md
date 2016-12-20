# go-pro

Practical examples for a introductory presentation about go programming

## Requirements

 * Go 1.6 or newer

## How to run the samples

 * Simply execute: `go run <path-to-the-sample>/main.go`

Example: `go run 01-helloworld/main.go`

## Build the example in a executable binary

 * Run bo build: `go build -o <binary-name> -i <path-to-the-sample>/main.go`

Example:
```shell
go build -o hello 01-helloworld/main.go
./hello # outputs "Hello World!"
```
