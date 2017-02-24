package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/bigomby/hashcode2017/internal/input"
	"github.com/bigomby/hashcode2017/internal/types"
)

// Videos is cool
type Videos []*types.Video

func (slice Videos) Len() int {
	return len(slice)
}

func (slice Videos) Less(i, j int) bool {
	return slice[i].Size < slice[j].Size
}

func (slice Videos) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

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

	// Sort videos
	var videos Videos
	endpoints, videos, servers, _ := input.ParseHeader(f)
	sort.Sort(videos)

	// Sort every endpoint
	for _, endpoint := range endpoints {
		var sortable types.Connections
		sortable = endpoint.Connections
		sort.Sort(sortable)
	}

	for _, video := range videos {
	serverLoop:
		for _, server := range servers {
			if server.Capacity >= video.Size {
				server.Videos = append(server.Videos, video)
				server.Capacity -= video.Size
				break serverLoop
			}
		}
	}

	// The output
	var usedServersCount int
	for _, server := range servers {
		if len(server.Videos) > 0 {
			usedServersCount++
		}
	}

	output := new(bytes.Buffer)
	fmt.Fprintf(output, "%d\n", usedServersCount)

	for i, server := range servers {
		fmt.Fprintf(output, "%d", i)
		for _, video := range server.Videos {
			fmt.Fprintf(output, " %d", video.ID)
		}

		fmt.Fprintf(output, "\n")
	}

	outfile, err := os.Create("output")
	outfile.Write(output.Bytes())
}
