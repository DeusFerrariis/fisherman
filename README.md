![GitHub issues](https://img.shields.io/github/issues/DeusFerrariis/fisherman?style=flat-square)
![GitHub License](https://img.shields.io/github/license/DeusFerrariis/fisherman?style=flat-square)

# 🎣 fisherman
Your CLI tool to execute commands on file changes

Fisherman is inspired by similar tools and is a "Re-write in X" project.
It is very much W.I.P. and should not be used in critical environments or such where the file system
is not backed up or versioned.

# Installation

## Manual

### Requirements

- Go >= 1.21.3

### Steps

1. Clone the repository locally to compile

`$ git clone https://github.com/DeusFerrariis/fisherman.git`

2. Compile/build application

`$ go build .`

3. Add binary to your path

<span style="font-weight: bold; color: red;">INFO</span> | I don't know if this works for Windows; if not create an issue! 

``` 
$ fisherman --help
NAME:
   fisherman - watch a file or directory and then execute a command after

USAGE:
   fisherman [global options] command [command options] [arguments...]

COMMANDS:
   watch    
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

```
