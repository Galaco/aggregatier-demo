package main

import (
	"github.com/galaco/aggregatier/cmd/smite/common"
	"github.com/galaco/aggregatier/models"
)

func main() {
	api := common.NewApi("devId", "authKey")

	session,err := api.CreateSession()
	if err != nil {
		panic(err)
	}
	heroes,err := api.GetGods(session.Id)
	if err != nil {
		panic(err)
	}

	models.InitDB("root:password@tcp(localhost:3306)/aggregatier")

	for _,hero := range heroes {
		if err := models.InsertHero(hero.Id, hero.Name, 1, hero.IconUrl); err != nil {
			panic(err)
		}
	}
}
