# base64


[![Go Report Card](https://goreportcard.com/badge/github.com/nikhilsbhat/base64)](https://goreportcard.com/report/github.com/nikhilsbhat/base64)

A utility that helps to encode and decode your data in [base64](https://en.wikipedia.org/wiki/Base64).

## Introduction

There are multiple online platform where one can easily encode/ decode their data into base64 format.
What if someone wants to achieve the same without connecting to internet?

Yes, exactly in such circumstances this utility base64 comes in handy.

## Installation

```golang
go get -u github.com/nikhilsbhat/base64
go build
```
One can start using the executable just like any other go application.

If incase few need to use this in your piece of code? import package in your code.
```golang
import (
    "github.com/nikhilsbhat/base64"
)
```

### base64 commands

```bash
base64 [command] [flags]
```
Make sure appropriate command is used for the actions, to check the available commands and flags use `base64 --help`

```bash
This will help user to encode and decode data in base64.


Usage:
  base64 [command] [flags]

Available Commands:
  decode      command to decode the string
  encode      command to encode the string
  help        Help about any command

Flags:
  -h, --help            help for base64
  -s, --string string   string which has to be encoded or decoded

Use "base64 [command] --help" for more information about a command."
```

### `base64 encode`

To encode the data you with in base64, command `encode` will be helpful.

```bash
base64 encode -s "content to be encoded"
```
