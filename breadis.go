package main

import (
	"github.com/inominate/breadis/srv"
)

func main() {
	srv.Listen()
	select {}
}
