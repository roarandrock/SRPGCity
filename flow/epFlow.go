package flow

import (
	"SRPGCity/action"
	"SRPGCity/inputs"
	"SRPGCity/model"
	"fmt"
)

//Cheating, for basic episode testing
func IntroEp() {
	fmt.Println("It's dangerous in here. Clear out the Monsters with your buddy.")
}

func MissionEp() {
	//need 2 monsters, two characters and commands
	mCont := true
	mOptions := []string{"Move", "Fight", "Quit"} //check notes
	rCount := 1
	//each round will draw the map and ask the player to do things
	for mCont == true {
		fmt.Println("Round", rCount)
		models.MakeMap() //this draws things, has the rooms
		r1 := inputs.StringarrayInput(mOptions)
		switch r1 {
		case 1:
			action.MoveEp()
		case 2:
			action.FightEp()
		case 3:
			mCont = false
		default:
			fmt.Println("That's no option")
			mCont = false
		}
		rCount++
	}
}
