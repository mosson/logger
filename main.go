package main

import (
	"fmt"
	"logger/log"
	"os"
	"sync/atomic"
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
	stop := make(chan bool)
	var op uint64

	for i := 0; i < 100; i++ {
		go func(i int) {
			l3.Error <- fmt.Sprintf("Hello, %d", i)
			l3.Info <- fmt.Sprintf("Hello, %d", i)
			atomic.AddUint64(&op, 1)

			if atomic.LoadUint64(&op) == uint64(100) {
				stop <- true
			}
		}(i)
	}

	<-stop

}
