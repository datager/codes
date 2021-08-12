package main

import (
	"codes/gocodes/dg/controllers/language"
	"codes/gocodes/dg/utils/signal"
)

func main() {
	language.TickerLog()
	signal.WaitForExit()
}
