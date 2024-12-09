package main

import (
	"server/connection"
	runproto "server/run_proto"
)

func main() {

	connection.InitDb()

	runproto.RunnerProtos()

}
