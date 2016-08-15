package log

import (
	"io/ioutil"
	"runtime"
	"sync/atomic"
	"testing"
)

func BenchmarkStaticLogger(b *testing.B) {
	l := NewLogger(DEBUG, ioutil.Discard)

	for i := 0; i < b.N; i++ {
		l.Info("Hello, World")
	}
}

func BenchmarkAsyncLogger(b *testing.B) {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	l := NewAsyncLogger(DEBUG, ioutil.Discard)
	go l.Run()
	defer l.Stop()

	var ops uint64
	length := b.N
	exit := make(chan bool)

	for i := 0; i < length; i++ {
		go func(i int) {
			l.Info <- "Hello, World"
			atomic.AddUint64(&ops, 1)
			if atomic.LoadUint64(&ops) == uint64(length) {
				exit <- true
			}
		}(i)
	}

	<-exit

}
