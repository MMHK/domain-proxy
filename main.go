// svnexport project main.go
package main

import (
	"domain-proxy/lib"
	"flag"
	"log"
	"os"
	"runtime"
)

func main() {
	confPath := flag.String("c", "config.json", "config file")
	flag.Parse()
	config, err := lib.NewConfig(*confPath)
	if err != nil {
		log.Println(err)
		os.Exit(404)
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	s := lib.NewService(config.Listen, config)

	log.Println("Domain service running, listen on:", config.Listen)

	s.Start()
}
