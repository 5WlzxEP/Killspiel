package ResultCollector

import (
	"Killspiel/pkg/database"
	"log"
	"strings"
)

type FinishedGames struct {
	GameId  int64
	Correct float64
}

func CheckUnfinishedGames() (finishedGames []FinishedGames) {
	rows, err := database.GetUnfinishedGames.Query()
	if err != nil {
		return nil
	}

	prefixes := map[string]*collector{}
	for _, v := range collectors {
		prefixes[v.Prefix()] = v
	}

	var id int64
	var info string
	for rows.Next() {
		err = rows.Scan(&id, &info)
		if err != nil {
			log.Println(err)
			continue
		}

		sub := strings.SplitN(info, ",", 2)
		if len(sub) < 1 {
			continue
		}

		if coll, ok := prefixes[sub[0]]; ok {
			result, err := coll.Resolve(info)
			if err != nil {
				continue
			}
			finishedGames = append(finishedGames, FinishedGames{id, result})
		} else {
			log.Printf("Unknown prefix for id %d: %s\n", id, sub[0])
		}
	}

	return
}
