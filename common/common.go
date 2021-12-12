package common

import (
	"flag"
)

var(
	port = flag.Uint64("port", 9000, "Port for communication")
)

func Port() uint64 {
	return *port
}