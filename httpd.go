package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mas9612/httpd/http"
)

type options struct {
	Port int `short:"p" long:"port" default:"8080" description:"Listen port"`
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	server := &http.Server{
		Port: opts.Port,
	}
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
