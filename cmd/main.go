package main

import (
	"covid/config"
	"covid/infrastructures"
	"covid/internal/adapters/covidadt"
	"covid/internal/core/service/covidsvc"
	"covid/internal/handlers"
	"covid/internal/repository"
)

func main() {
	config.Init()
	// infrastructure
	mc := infrastructures.NewMongoDB()
	// repository
	cr := repository.NewCovidRepo(mc, config.Get().Mongo.Database)
	// adapter
	ca := covidadt.New()

	csvc := covidsvc.New(ca, cr)

	chdl := handlers.NewCovidHdl(csvc)

	if err := chdl.Start(); err != nil {
		panic(err)
	}
}
