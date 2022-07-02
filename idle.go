package gobotapi

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Idle() {
	waitStart := make(chan os.Signal)
	signal.Notify(
		waitStart,
		syscall.SIGINT,
	)
	go func() {
		<-waitStart
		os.Exit(0)
	}()
	for {
		time.Sleep(10 * time.Second)
	}
}
