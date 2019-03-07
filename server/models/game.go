package models

import (
	"fmt"
)

type Game struct {
	Id        int
	Name      string
	ShortName string
}

func FindGame(id int) (*Game, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM games WHERE id = %d", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	game := new(Game)
	rows.Next()
	err = rows.Scan(&game.Id, &game.Name, &game.ShortName)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func AllGames() ([]*Game, error) {
	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := make([]*Game, 0)
	for rows.Next() {
		game := new(Game)
		err := rows.Scan(&game.Id, &game.Name, &game.ShortName)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return games, nil
}
