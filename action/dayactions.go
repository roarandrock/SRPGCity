package action

import (
	"SRPGCity/inputs"
	"SRPGCity/model"
	"fmt"
)

func MoveEp() {
	cP := models.GetCurrentPlayer()
	north := "North"
	south := "South"
	east := "East"
	west := "West"
	//look at current location, give options to move
	mOptionsStr := []string{} //need collision for walls and monsters
	mOptionsInt := models.CollisionCheck(cP.Coord, cP.RoomId)
	for _, v := range mOptionsInt {
		switch v {
		case 1:
			mOptionsStr = append(mOptionsStr, north)
		case 2:
			mOptionsStr = append(mOptionsStr, south)
		case 3:
			mOptionsStr = append(mOptionsStr, west)
		case 4:
			mOptionsStr = append(mOptionsStr, east)
		}
	}
	if len(mOptionsStr) == 0 {
		mOptionsStr = append(mOptionsStr, "Trapped and")
	}
	r1 := inputs.StringarrayInput(mOptionsStr)
	mDir := mOptionsStr[r1-1]
	switch mDir {
	case north:
		cP.Coord[1]++
	case south:
		cP.Coord[1]--
	case east:
		cP.Coord[0]++
	case west:
		cP.Coord[0]--
	default:
		fmt.Println("Nowhere to go")
	}
	models.UpdateCurrentPlayer(cP)
}

func FightEp() {
	//choose Attack
	cP := models.GetCurrentPlayer()
	fmt.Println("Melee:\n", cP.CurJorb.AttackMelee.Name, ":", cP.CurJorb.AttackMelee.Descr)
	options := []string{"Melee", "Range"}
	r1 := inputs.StringarrayInput(options)
	switch r1 {
	case 1:
		//get monsters in range
		attkR := 1 //range of Attack
		rangeEp(cP, attkR)
		//perform attach
	case 2:
		fmt.Println("Not now cowboy")
	}
}

func rangeEp(cP models.Player, attkR int) {
	mRay := models.RangeCheck(cP.Coord, cP.RoomId, attkR) //returns directions for monsters
	//look at collission logic, but need find monster by coord still
	fmt.Println(mRay)
}
