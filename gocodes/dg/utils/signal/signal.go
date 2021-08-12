package signal

import (
	"os"
	"os/signal"
	"syscall"
)

////////////////// signal /////////////////////
func WaitForExit() os.Signal {
	return WaitForSignal(syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
}

func WaitForSignal(sources ...os.Signal) os.Signal {
	var s = make(chan os.Signal, 1)
	defer signal.Stop(s) //the second Ctrl+C will force shutdown

	signal.Notify(s, sources...)
	return <-s //blocked
}
