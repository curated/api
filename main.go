package main

import (
	"flag"

	"github.com/curated/elastic/config"
	"github.com/curated/elastic/server"
)

const root = "./"

func main() {
	flag.Parse()
	c := config.New(root)
	server.New(c).Start()
}
