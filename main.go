//go:generate statik -src=./gui
package main

import (
	"github.com/dtchanpura/pinboard/server"
)

func main() {
	server.StartListener()
}
