package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/bigomby/hashcode2017/internal/types"

	log "github.com/Sirupsen/logrus"
)

// Logger logs things
var Logger *log.Logger

// ParseHeader parses the fist line of the file
func ParseHeader(input io.Reader) ([]types.Endpoint, []*types.Video, []*types.CacheServer, []types.RequestDescription) {
	var videosCount int
	var endpointsCount int
	var requestsDescriptionsCount int
	var cacheServersCount int
	var capacity int

	endpoints := make([]types.Endpoint, 1)
	requests := make([]types.RequestDescription, 1)

	scanner := bufio.NewScanner(input)
	if err := scanner.Err(); err != nil {
		Logger.Fatal(err)
	}

	scanner.Scan()
	header := scanner.Text()

	fmt.Sscanf(header, "%d %d %d %d %d",
		&videosCount,
		&endpointsCount,
		&requestsDescriptionsCount,
		&cacheServersCount,
		&capacity,
	)

	serverTemplate := types.CacheServer{
		Capacity: capacity,
	}

	var servers []*types.CacheServer
	for i := 0; i < cacheServersCount; i++ {
		serverTemplateCopy := serverTemplate
		servers = append(servers, &serverTemplateCopy)
	}

	for _, server := range servers {
		server = new(types.CacheServer)
		server.Capacity = capacity
	}

	scanner.Scan()
	videos := parseVideos(scanner.Text())
	parseVideos(scanner.Text())

	for i := 0; i < endpointsCount; i++ {
		endpoints = append(endpoints, parseEndpoint(scanner, servers))
	}

	for i := 0; i < requestsDescriptionsCount; i++ {
		requests = append(requests, parseRequest(scanner, videos, endpoints))
	}

	return endpoints, videos, servers, requests
}

func parseVideos(line string) []*types.Video {
	var videos []*types.Video
	sizes := strings.Split(line, " ")

	for i, size := range sizes {
		sizeInt, _ := strconv.ParseInt(size, 10, 32)
		videos = append(videos, &types.Video{
			ID:   i,
			Size: int(sizeInt),
		})
	}

	return videos
}

func parseEndpoint(scanner *bufio.Scanner, servers []*types.CacheServer) types.Endpoint {
	var connections int
	var latency int

	scanner.Scan()
	header := scanner.Text()
	fmt.Sscanf(header, "%d %d", &latency, &connections)

	endpoint := types.Endpoint{
		Latency:     latency,
		Connections: make([]types.Connection, 0),
	}

	for i := 0; i < connections; i++ {
		connection := types.Connection{}
		var connectionID int

		scanner.Scan()
		attributes := scanner.Text()
		fmt.Sscanf(attributes, "%d %d", &connectionID, &connection.CacheLatency)
		connection.Server = servers[connectionID]
		endpoint.Connections = append(endpoint.Connections, connection)
	}

	return endpoint
}

func parseRequest(
	scanner *bufio.Scanner,
	videos []*types.Video,
	endpoints []types.Endpoint,
) types.RequestDescription {
	var videoID int
	var endpointID int
	var amount int

	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d %d", &videoID, &endpointID, &amount)
	// Logger.Debugf("[REQUEST] Amount: %d | Video: %d | Endpoint: %d", amount, videoID, endpointID)

	return types.RequestDescription{
		Amount: amount,
		Video:  videos[videoID],
		Source: &endpoints[endpointID],
	}
}
