package models

import "fmt"

/*
Jorbs
*/

//Jorb is base struct
type Jorb struct {
	CompanyName string
	Industry    string
	Title       string
	Schedule    []int
	Minstats    Jorbmin
	Impact      Jorbimpact
}

//Jorbmin minimum stats needed
type Jorbmin struct {
	MinStats  PlayerStats
	MinTraits PlayerTraits
}

//Jorbimpact impact on stats from taking/working jorb
type Jorbimpact struct {
	ImpactStats  PlayerStats
	ImpactTraits PlayerTraits
	ImpactMeters PlayerMeters
}

//FirstJorb first job possible
var FirstJorb Jorb
var SecondJorb Jorb
var NoJorb Jorb
var JorbBoard = []Jorb{}

func JorbGet(sJ Jorb) Jorb {
	fJ := FirstJorb
	for _, v := range JorbBoard {
		if v.Title == sJ.Title && v.CompanyName == sJ.CompanyName {
			fJ = v
			//fmt.Println("test2:", sJ, fJ)
		}
	}
	return fJ
}

func GetCurJorb() Jorb {
	cp := GetCurrentPlayer()
	return cp.CurJorb
}

//JorbList returns an array of Jorbs,removing "Nothing" ones
func JorbList() (int, []Jorb) {
	i := 0
	jList := []Jorb{}
	for _, v := range JorbBoard {
		if v.Industry != "Nothing" {
			jList = append(jList, v)
			i++
		}
	}
	return i, jList
}

func JorbDisplay(j1 Jorb) {
	fmt.Printf("%s: %s\n", "Company", j1.CompanyName)
	fmt.Printf("%s: %s\n", "Title", j1.Title)
	fmt.Printf("%s: %v\n", "Income", j1.Impact.ImpactTraits.Income)
}

//SetJorb sets initial job stats
func SetJorb() {
	FirstJorb.CompanyName = "Fun First Jorb Corp"
	FirstJorb.Industry = "Software Testing"
	FirstJorb.Title = "Entry Level"
	FirstJorb.Schedule = []int{0, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1} //busy morning, working, then 1s

	//Zero for all mins on first jorb
	//ImpactMeters
	FirstJorb.Impact.ImpactMeters.Energy = -50
	FirstJorb.Impact.ImpactMeters.Hygiene = -10
	FirstJorb.Impact.ImpactMeters.Soul = -1
	FirstJorb.Impact.ImpactMeters.Stress = 20
	//ImpactStats
	FirstJorb.Impact.ImpactStats.Connections = 5
	FirstJorb.Impact.ImpactStats.Experience = 10
	FirstJorb.Impact.ImpactStats.Knowledge = 5
	//ImpactTraits
	FirstJorb.Impact.ImpactTraits.Income = 20
	FirstJorb.Impact.ImpactTraits.Independence = 40

	//second one
	SecondJorb.CompanyName = "Serious Corp"
	SecondJorb.Industry = "Software Testing"
	SecondJorb.Title = "Professional Level"
	SecondJorb.Schedule = []int{0, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1} //busy morning, working, then 1s

	//MinStats
	SecondJorb.Minstats.MinStats.Connections = 10
	SecondJorb.Minstats.MinStats.Experience = 20
	SecondJorb.Minstats.MinStats.Knowledge = 10
	//MinTraits
	SecondJorb.Minstats.MinTraits.Education = 20
	SecondJorb.Minstats.MinTraits.Independence = 30
	SecondJorb.Minstats.MinTraits.Lucky = 10
	SecondJorb.Minstats.MinTraits.Smooth = 10
	//ImpactMeters
	SecondJorb.Impact.ImpactMeters.Energy = -30
	SecondJorb.Impact.ImpactMeters.Hygiene = -20
	SecondJorb.Impact.ImpactMeters.Soul = -5
	SecondJorb.Impact.ImpactMeters.Stress = 40
	//ImpactStats
	SecondJorb.Impact.ImpactStats.Connections = 10
	SecondJorb.Impact.ImpactStats.Experience = 10
	SecondJorb.Impact.ImpactStats.Knowledge = 10
	//ImpactTraits
	SecondJorb.Impact.ImpactTraits.Income = 40
	SecondJorb.Impact.ImpactTraits.Independence = 30

	//Unemployed
	NoJorb.CompanyName = "Nowhere"
	NoJorb.Industry = "Nothing"
	NoJorb.Title = "Unemployed"
	NoJorb.Schedule = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1} //busy morning, working, then 1s

	//MinStats
	NoJorb.Minstats.MinStats.Connections = 0
	NoJorb.Minstats.MinStats.Experience = 0
	NoJorb.Minstats.MinStats.Knowledge = 0
	//MinTraits
	NoJorb.Minstats.MinTraits.Education = 0
	NoJorb.Minstats.MinTraits.Independence = 0
	NoJorb.Minstats.MinTraits.Lucky = 0
	NoJorb.Minstats.MinTraits.Smooth = 0
	//ImpactMeters
	NoJorb.Impact.ImpactMeters.Energy = 0
	NoJorb.Impact.ImpactMeters.Hygiene = 0
	NoJorb.Impact.ImpactMeters.Soul = 0
	NoJorb.Impact.ImpactMeters.Stress = 0
	//ImpactStats
	NoJorb.Impact.ImpactStats.Connections = 5
	NoJorb.Impact.ImpactStats.Experience = 5
	NoJorb.Impact.ImpactStats.Knowledge = 5
	//ImpactTraits
	NoJorb.Impact.ImpactTraits.Income = 1
	NoJorb.Impact.ImpactTraits.Independence = 100

	//initial Jorb
	cp := GetCurrentPlayer()
	cp.CurJorb = FirstJorb
	UpdateCurrentPlayer(cp)

	//set JorbBoard
	JorbBoard = append(JorbBoard, FirstJorb)
	JorbBoard = append(JorbBoard, SecondJorb)
	JorbBoard = append(JorbBoard, NoJorb)

}
