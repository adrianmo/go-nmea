// Package main is the netflix switcher service.
package main

import (
	"flag"
	"flixer"
	"fmt"
	"log"
	"os"
)

var (
	addr    = flag.String("address", "", "Address to listen on")
	port    = flag.Int("port", 8080, "Port to listen on")
	logfile = flag.String("logfile", "", "Logfile to write")
)

func openLog() (*os.File, error) {
	var err error
	var f *os.File
	if _, err := os.Stat(*logfile); os.IsNotExist(err) {
		f, err = os.Create(*logfile)
	} else {
		f, err = os.OpenFile(*logfile, os.O_APPEND|os.O_WRONLY, 0600)
	}
	return f, err
}

func main() {
	flag.Parse()
	if *logfile != "" {
		f, err := openLog()
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}
	if err := flixer.Start(*addr, *port); err != nil {
		fmt.Printf("Error starting flixer: %v\n", err)
	}
}
