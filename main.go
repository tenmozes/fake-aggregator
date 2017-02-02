package main

import (
	"flag"
	"log"
	"time"

	"github.com/tenmozes/fake-aggregator/aggregator"
	"github.com/tenmozes/fake-aggregator/server"
)

var (
	port     = flag.Int("port", 8080, "server port")
	delay    = flag.Int("delay", 400, "max aggregator handlers delay in ms")
	deadline = flag.Int("deadline", 500, "max response time from aggregator handles in ms")
	factor   = flag.Int("factor", 50, "max number limits in aggregator handlers")
)

func init() {
	log.SetPrefix("[FakeAggregator] ")
	flag.Parse()
}

func main() {
	s := server.NewServer(aggregator.NewMapper(),
		server.WithDelay(*delay),
		server.WithDeadline(time.Duration(*deadline)*time.Millisecond),
		server.WithRandomFactor(*factor),
	)
	log.Fatal(s.Run(*port))
}
