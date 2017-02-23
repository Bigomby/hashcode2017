package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/bigomby/hashcode2017/internal/input"
)

func main() {
	input.Logger = log.New()

	var debug = flag.Bool("debug", false, "Debug info")
	var filename = flag.String("filename", "", "Filename")

	flag.Parse()

	if filename == nil {
		flag.Usage()
		os.Exit(1)
	}
	if *debug {
		input.Logger.Level = log.DebugLevel
	}

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	input.ParseHeader(f)
}
