package signal

import (
	"os"
	"os/signal"
	"syscall"
)

//CheckSySignal --
func CheckSySignal(fc func(string)) {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fc("Server Exiting..")
	os.Exit(0)
}

// CheckSySignalWithFunc --
func CheckSySignalWithFunc(f func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	f()
	os.Exit(0)
}
