package simple_pbft

import (
	"os"

	"github.com/dylenfu/pbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
