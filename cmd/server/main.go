package main

import (
	"fmt"
	"os"

	"github.com/rrgmc/debefix-sample-app/pkg/app"
	"github.com/rrgmc/debefix-sample-app/pkg/config"
)

func main() {
	cfg, err := config.LoadFile(os.Getenv("SERVER_CONFIG_FILE"))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error loading config file: %s", err)
		os.Exit(1)
	}

	serverApp := app.NewApp(cfg)
	err = serverApp.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error running application: %s", err)
		os.Exit(1)
	}
}
