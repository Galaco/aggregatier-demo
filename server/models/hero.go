package models

import "fmt"

type Hero struct {
	Id      int
	Name    string
	GameId  int
	IconUrl string
}

func InsertHero(id int, name string, gameId int, iconUrl string) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO heroes (id, name, game_id, icon_url) VALUES (%d, \"%s\", %d, \"%s\")", id, name, gameId, iconUrl))
	return err
}

func AllHeroes(gameId int) ([]*Hero, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM heroes WHERE game_id = %d", gameId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	heroes := make([]*Hero, 0)
	for rows.Next() {
		hero := new(Hero)
		err := rows.Scan(&hero.Id, &hero.Name, &hero.GameId, &hero.IconUrl)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, hero)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return heroes, nil
}
