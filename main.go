package main

import (
	"log"

	"gitlab.com/gocastsian/writino/app"
	"gitlab.com/gocastsian/writino/config"
	v1 "gitlab.com/gocastsian/writino/delivery/http/v1"
)

func main() {

	cfg, err := config.LoadCfg("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(v1.New(app, cfg.Server).Run())
}
