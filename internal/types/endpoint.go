package types

// Connections does stuff
type Connections []Connection

func (slice Connections) Len() int {
	return len(slice)
}

func (slice Connections) Less(i, j int) bool {
	return slice[i].CacheLatency < slice[j].CacheLatency
}

func (slice Connections) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Endpoint represents a group of users connecting to the Internet in the same
// geographical area (for example, a neighborhood in a city). Every endpoint is
// connected to the data center. Additionally, each endpoint may (but doesn’t
// have to) be connected to 1 or more cache servers​.
//
// Each endpoint is characterized by the latency of its connection to the data
// center (how long it takes to serve a video from the data center to a user in
// this endpoint), and by the latencies to each cache server that the endpoint
// is connected to (how long it takes to serve a video stored in the given cache
// server to a user in this endpoint).
type Endpoint struct {
	Connections []Connection
	Latency     int
}

// Connection is a connection from an endpoint to a Cache Server
type Connection struct {
	Server       *CacheServer
	CacheLatency int
}
