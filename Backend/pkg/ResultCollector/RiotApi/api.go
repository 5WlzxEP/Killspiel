package RiotApi

import (
	"fmt"
)

type Summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (a Api) getSummonerByName(name string) (*Summoner, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", a.Server, name)

	s, err := get[Summoner](a, url)
	return s, err
}
