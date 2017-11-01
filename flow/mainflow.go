package flow

import (
	"SRPGCity/check"
	"SRPGCity/inputs"
	"SRPGCity/model"
	"fmt"
	"strconv"
)

//MainFlow runs the game loop
func MainFlow() error {
	fmt.Println("Welcome to the jungle.")
	cont, err := starterFlow()
	check.ErrorCheck(err)
	for cont == true {
		//game loop
		_, err = gameLoopFlow() //should put cont here?
		//check to continue
		cont = check.GameEndCheck()
	}
	endFlow()
	return err
}

func starterFlow() (bool, error) {
	var cont bool
	_, err := strconv.Atoi("-42") //error cheat
	fmt.Println("Do you want to play a game?")
	r1 := inputs.StringarrayInput([]string{"Yes", "No"})
	if r1 == 1 {
		cont = true
	} else {
		cont = false
		return cont, err
	}
	//narrative.Intro()
	models.SetPlayerAll()
	models.SetMon()
	models.SetJorb() //setting up job board
	//models.PlayerJorbUpdate() //setting first job impact and schedule
	//setting up city
	models.SetCity()
	models.SetEvents()
	//setting Time
	models.InitTime()
	return cont, err
}

func gameLoopFlow() (bool, error) {
	cont := true
	_, err := strconv.Atoi("-42") //error cheat
	for cont == true {
		dayFlowLoop() //episode loop
		cont = check.GameEndCheck()
	}
	return cont, err
}

func dayFlowLoop() {
	models.ShowTime() //test
	//apartment load
	//leave apartment, start episode
	//Intro
	IntroEp()
	//mission
	MissionEp()
	//outro
	endOfWeek() //update episode, apartment
	//loop
}

//move to checker/events
func endOfWeek() {
	ot := models.GetTime()
	ot.EpN++
	models.UpdateTime(ot)
}

func endFlow() {
	fmt.Println("It's all over.")
}
