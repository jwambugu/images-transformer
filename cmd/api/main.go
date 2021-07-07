package main

import (
	"fmt"
	"github.com/jwambugu/images-transformer/pkg/config"
	"log"
	"os"
	"os/signal"
)

var (
	configKeys     *config.Config
	absolutePath   = config.GetAbsolutePath()
	configKeysFile = fmt.Sprintf("%s%s", absolutePath, ".keys.json")
	storagePath    = fmt.Sprintf("%s%s", absolutePath, "storage")
)

type application struct {
	config *config.Config
}

func init() {
	var err error

	configKeys, err = config.Read(configKeysFile)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &application{
		config: configKeys,
	}

	fmt.Println(storagePath)

	fiberApp := app.routes()
	addr := fmt.Sprintf(":%d", app.config.AppPort)

	serverShutdownChan := make(chan os.Signal, 1)
	signal.Notify(serverShutdownChan, os.Interrupt)
	signal.Notify(serverShutdownChan, os.Kill)

	go func() {
		<-serverShutdownChan
		log.Println("Gracefully shutting down the web server...")

		_ = fiberApp.Shutdown()
	}()

	if err := fiberApp.Listen(addr); err != nil {
		log.Panic(err)
	}
}
