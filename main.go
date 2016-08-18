package main

import (
	"flag"
	"fmt"

	"github.com/cenan/mergen/web"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	port := flag.Int("port", 7777, "server port")
	flag.Parse()
	open.Run(fmt.Sprintf("http://localhost:%d", *port))
	web.StartServer(*port)
}
