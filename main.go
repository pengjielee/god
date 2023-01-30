package main

import (
	"flag"
	"fmt"
	"time"
)

const HELPINFO = `
NAME:
   god - god tool by go

USAGE:
   god [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   list       list go versions
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
`

const VERSION = "0.0.1"

func main() {
	var h bool
	flag.BoolVar(&h, "h", false, "")
	flag.BoolVar(&h, "help", false, "show help")

	var v bool
	flag.BoolVar(&v, "v", false, "")
	flag.BoolVar(&v, "version", false, "show version")

	var d bool
	flag.BoolVar(&d, "d", false, "")
	flag.BoolVar(&d, "date", false, "show current date")

	var n bool
	flag.BoolVar(&n, "n", false, "")
	flag.BoolVar(&n, "now", false, "show current timestamp")

	flag.Parse()

	if h {
		fmt.Println(HELPINFO)
	}

	if v {
		fmt.Println(VERSION)
	}

	if d {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	}

	if n {
		fmt.Println(time.Now().Unix())
	}
}
