package cfn

import (
	"encoding/json"
	"fmt"
)

type Battle struct {
	Ratio  float32 `json:"ratio"`
	Total  int     `json:"total"`
	Wins   int     `json:"wins"`
	Losses int     `json:"losses"`
}

type Character struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Serialized struct {
	Id              string    `json:"id"`
	Region          string    `json:"region"`
	Device          string    `json:"device"`
	Status          string    `json:"status"`
	Title           string    `json:"title"`
	Dojo            string    `json:"dojo"`
	Twitter         string    `json:"twitter"`
	Level           int       `json:"level"`
	AccountAge      string    `json:"account_age"`
	TotalMatches    int       `json:"total_matches"`
	Character       Character `json:"character"`
	League          string    `json:"league"`
	LeaguePoints    int       `json:"league_points"`
	Ranking         int       `json:"ranking"`
	FirstMatch      string    `json:"first_match"`
	MostRecentMatch string    `json:"most_recent_match"`
	Ranked          Battle    `json:"ranked"`
	Casual          Battle    `json:"casual"`
	Lounge          Battle    `json:"lounge"`
}

func (p Profile) ToJSON() (response []byte) {
	serialized := Serialized{
		Id:           p.GetFighterId(),
		Region:       p.GetRegion(),
		Device:       p.GetDevice(),
		Status:       p.GetStatus(),
		Title:        p.GetTitle(),
		Dojo:         p.GetDojo(),
		Twitter:      p.GetTwitter(),
		Level:        p.GetPlayerLevel(),
		AccountAge:   p.GetAccountAge(),
		TotalMatches: p.GetTotalMatches(),
		Character: Character{
			Name:  p.GetCharacterName(),
			Level: p.GetCharacterLevel(),
		},
		League:          p.GetLeague(),
		LeaguePoints:    p.GetLeaguePoints(),
		Ranking:         p.GetRanking(),
		FirstMatch:      p.GetFirstMatch(),
		MostRecentMatch: p.GetMostRecentMatch(),
		Ranked: Battle{
			Ratio:  p.GetRankedMatchesRatio(),
			Total:  p.GetRankedMatchesTotal(),
			Wins:   p.GetRankedMatchesWins(),
			Losses: p.GetRankedMatchesLosses(),
		},
		Casual: Battle{
			Ratio:  p.GetCasualMatchesRatio(),
			Total:  p.GetCasualMatchesTotal(),
			Wins:   p.GetCasualMatchesWins(),
			Losses: p.GetCasualMatchesLosses(),
		},
		Lounge: Battle{
			Ratio:  p.GetBattleLoungeMatchesRatio(),
			Total:  p.GetBattleLoungeMatchesTotal(),
			Wins:   p.GetBattleLoungeMatchesWins(),
			Losses: p.GetBattleLoungeMatchesLosses(),
		},
	}

	response, err := json.Marshal(serialized)
	if err != nil {
		fmt.Println(err, serialized)
	}

	return
}
