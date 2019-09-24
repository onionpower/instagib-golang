package main

import (
	bc "instagob/basicconcurrency"
)

func main() {
	wait1 := make(chan bool)
	ch1 := bc.BoringGen("you", wait1)
	wait2 := make(chan bool)
	ch2 := bc.BoringGen("me", wait2)
	ch := bc.FanIn(ch1, ch2)
	bc.Listen(ch)
}
