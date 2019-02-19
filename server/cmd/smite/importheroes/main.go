package main

import (
	"github.com/galaco/aggregatier/cmd/smite/common"
	"github.com/galaco/aggregatier/models"
	"log"
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
		log.Println(hero.Id)
		if err := models.InsertHero(hero.Id, hero.Name, 1); err != nil {
			panic(err)
		}
	}
}
