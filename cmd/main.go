package main

import (
	"flag"
	"fmt"

	"snippetbox/cmd/web"
)

// main calls run that initializes a http server
func main() {
	port := flag.Int("port", 4000, "HTTP network address")
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	web.Run(addr)
}
