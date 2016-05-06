package cmd

import (
	"fmt"
)

func printProjectVersion() {
	fmt.Printf("PingPong %s (%s)\n", projectVersion, projectCommit)
}
