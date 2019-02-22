package models

import (
	"fmt"
)

type Tier struct {
	Id int
	Name string
	GameId string
}

func AllTiers(gameId int) ([]*Tier, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM tier WHERE game_id = %d", gameId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tiers := make([]*Tier, 0)
	for rows.Next() {
		tier := new(Tier)
		err := rows.Scan(&tier.Id, &tier.Name, &tier.GameId)
		if err != nil {
			return nil, err
		}
		tiers = append(tiers, tier)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tiers, nil
}