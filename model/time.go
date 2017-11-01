package models

import "fmt"

/*

And newer plan: TV episodes. Each mission is it's own episode. Less time issues

//season, episode, apartment or city

*/

type TimeMeter struct {
	SeasonN int
	EpN     int //epm structure? name phase
}

var currentTime = TimeMeter{}

func UpdateTime(dt TimeMeter) {
	currentTime = dt
}

func NextEpisode() {
	//make next ep
}

func GetTime() TimeMeter {
	ct := currentTime
	return ct
}

func InitTime() {
	nt := GetTime()
	nt.SeasonN = 1
	nt.EpN = 1
	UpdateTime(nt)
}

func ShowTime() {
	ct := GetTime()
	fmt.Println("Season", ct.SeasonN, "Episode", ct.EpN)
}
