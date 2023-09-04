package main

import (
	api "github.com/siddhantprateek/Trakd/api"
	config "github.com/siddhantprateek/Trakd/config"
)

func main() {

	svr := api.NewAPI(&config.APIConfiguration{})

	svr.Start()
}
