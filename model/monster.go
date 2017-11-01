package models

import "fmt"

/*
Basic monster models
Lots of cheats
Need to have default models but store specific instances in the episode (outside of models)
Will need an array of monsters to search through, maybe setup some separate groups of monsters to simplify this
*/

type Monster struct {
	Name     string
	Health   int
	MId      int
	RoomId   int
	Coord    []int
	Strategy int
	Attack   int
}

var Mon1 Monster
var Mon2 Monster

func GetMonById(MId int) Monster {
	var selectMon Monster
	if MId == 1 {
		selectMon = Mon1
	} else {
		selectMon = Mon2
	}
	return selectMon
}

//can only have a single monster on a spot
func GetMonByLoc(lCoord []int, rId int) Monster {
	var selectMon Monster
	mRay := []Monster{}
	mRay = append(mRay, Mon1, Mon2)
	for _, v := range mRay {
		if rId == v.RoomId {
			if v.Coord[0] == lCoord[0] {
				if v.Coord[1] == lCoord[1] {
					selectMon = v
				}
			}
		}
	}
	if selectMon.MId == 0 {
		fmt.Println("Cheat error, no monster found by location")
	}
	return selectMon
}

func UpdateMon(nMon Monster) {
	if nMon.MId == 1 {
		Mon1 = nMon
	} else {
		Mon2 = nMon
	}
}

//SetPlayerAll sets initial stats, assuming no jorb
func SetMon() {
	Mon1.Name = "Sloth"
	Mon1.Coord = []int{6, 1}
	Mon1.Health = 100
	Mon1.MId = 1
	Mon1.RoomId = 1
	//mon 2
	Mon2.Name = "Time"
	Mon2.Coord = []int{6, 8}
	Mon2.Health = 100
	Mon2.MId = 2
	Mon2.RoomId = 2
}
