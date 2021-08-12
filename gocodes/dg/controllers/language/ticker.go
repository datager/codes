package language

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func WrongTicker() {
	ch := make(chan string, 100)
	go func() {
		for {
			ch <- "y"
		}
	}()
	go func() {
		if err := http.ListenAndServe("127.0.0.1:9999", nil); err != nil {
			fmt.Println("http server fail")
		}
	}()
	for {
		select {
		case <-ch:
		case <-time.After(time.Minute * 3):
		}
	}
}

func RightTicker() {
	ch := make(chan string, 100)
	go func() {
		for {
			ch <- "y"
		}
	}()
	go func() {
		if err := http.ListenAndServe("127.0.0.1:9999", nil); err != nil {
			fmt.Println("http server fail")
		}
	}()
	tk := time.NewTimer(time.Minute * 3)
	for {
		select {
		case <-ch:
			if !tk.Stop() {
				<-tk.C
			}
		case <-tk.C:
		}
		tk.Reset(time.Minute * 3)
	}
	tk.Stop()
}

func TickerLog() {
	debugTimer := time.NewTimer(time.Duration(1 * time.Second))
	go func() {
		for {
			select {
			case <-debugTimer.C:
				logrus.Infof("[mem-plate] %v", 333)
			}
			debugTimer.Reset(1 * time.Second)
		}
		debugTimer.Stop()
	}()
}
