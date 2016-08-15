package main

import (
	"logger/log"
	"os"
)

func main() {
	l := log.DefaultLogger()
	l.Error("Hello, World")

	file, err := os.Create("sample.log")
	if err != nil {
		panic(err)
	}

	l2 := log.NewLogger(log.DEBUG, file)
	l2.Warn("Hip Hop")

	l3 := log.NewAsyncLogger(log.DEBUG, os.Stdout)
	go l3.Run()
	defer l3.Stop()

	l3.Error <- "Hello, AsyncLogger"
	l3.Info <- "Hello, AsyncLogger"
}
