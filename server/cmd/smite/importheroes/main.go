package main

import (
	"github.com/galaco/aggregatier/cmd/smite/common"
	"github.com/galaco/aggregatier/config"
	"github.com/galaco/aggregatier/models"
)

func main() {
	configuration, err := config.NewConfig("config.json")
	if err != nil {
		panic(err)
	}

	api := common.NewApi(configuration.Smite.DevId, configuration.Smite.AuthKey)

	session, err := api.CreateSession()
	if err != nil {
		panic(err)
	}
	heroes, err := api.GetGods(session.Id)
	if err != nil {
		panic(err)
	}

	models.InitDB(configuration.Database.Username +
		":" +
		configuration.Database.Password +
		"@tcp(" +
		configuration.Database.Host +
		")/" +
		configuration.Database.Name)

	for _, hero := range heroes {
		if err := models.InsertHero(hero.Id, hero.Name, 1, hero.IconUrl); err != nil {
			panic(err)
		}
	}
}
