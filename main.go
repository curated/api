package main

import (
	"flag"

	"github.com/curated/elastic/server"
)

func main() {
	flag.Parse()
	server.New().Start()
}
