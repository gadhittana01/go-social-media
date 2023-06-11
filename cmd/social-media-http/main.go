package main

import (
	"log"

	"github.com/gadhittana01/socialmedia/config"
	"github.com/gadhittana01/socialmedia/helper"
)

func main() {
	config := &config.GlobalConfig{}
	helper.LoadConfig(config)
	err := initApp(config)
	if err != nil {
		log.Println(err)
	}
}
