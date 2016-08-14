package main

import "logger/log"

func main() {
	l := log.DefaultLogger()
	l.Error("Hello, World")
}
