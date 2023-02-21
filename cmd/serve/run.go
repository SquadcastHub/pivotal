package serve

import (
	"fmt"
	"github.com/SquadcastHub/auxpkg/v2/store"
	"os"
	"pivotal/cmd/serve/config"
	"pivotal/cmd/serve/metrics"
)

func Run() {
	config.Init(os.Args[2:]...)
	metrics.Init()
	store.Initialize()
	fmt.Println("running server command")
}
