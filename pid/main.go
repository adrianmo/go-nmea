package main

import (
	"math/rand"
	"pid"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//pid.StartSimulation()
	rt := pid.RealTime{}
	rt.Begin("kettle")
}
