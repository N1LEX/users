package main

import (
	"flag"
	"log"
	"runtime"
)

var host, port string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&host, "-h", "127.0.0.1", "host")
	flag.StringVar(&port, "-p", ":9000", "port")
}

func main() {
	flag.Parse()

	// DB Connect and Apply Migrations
	InitDB()
	MakeMigrations()

	// APP Routing
	r := GetRouter()
	log.Fatal(r.Run())
}
