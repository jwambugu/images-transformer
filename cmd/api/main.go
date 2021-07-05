package main

import (
	"fmt"
	"gitlab.com/jwambugu/images_transformer/pkg/config"
	"log"
)

var (
	configKeys     *config.Config
	absolutePath   = config.GetAbsolutePath()
	configKeysFile = fmt.Sprintf("%s%s", absolutePath, ".keys.json")
)

func init() {
	var err error

	configKeys, err = config.Read(configKeysFile)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(absolutePath, configKeys)
}
