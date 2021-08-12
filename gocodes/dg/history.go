package main

import (
	"codes/gocodes/dg/utils/signal"
	"github.com/golang/glog"
)

func main() {



	sig := signal.WaitForExit()
	glog.Infof("Got signal: %v, exit", sig)
}
