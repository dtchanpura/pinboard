//go:generate statik -m -src=./gui
package main

import (
	"github.com/dtchanpura/pinboard/server"
)

func main() {
	server.StartListener()
}
