package serve

import (
	"fmt"
	"pivotal/pkg/config"
)

func Run() {
	config.Init()
	fmt.Println("running server command")
}
