package main

import (
	"flag"

	"github.com/cenan/mergen/web"
)

func main() {
	port := flag.Int("port", 7777, "server port")
	flag.Parse()
	web.StartServer(*port)
}
