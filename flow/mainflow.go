package flow

import (
	"UrbanRPG1/action"
	"UrbanRPG1/check"
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"UrbanRPG1/narrative"
	"fmt"
	"strconv"
)

//MainFlow runs the game loop
func MainFlow() error {
	fmt.Println("Welcome to the jungle")
	cont, err := starterFlow()
	check.ErrorCheck(err)
	for cont == true {
		//game loop
		_, err = gameLoopFlow()
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
	narrative.Intro()
	models.SetPlayerAll()
	models.SetJorb()          //setting up job board
	models.PlayerJorbUpdate() //setting first job impact and schedule
	//setting up city
	models.SetCity()
	models.SetEvents()
	//new flow
	//setting Time
	models.InitTime()
	return cont, err
}

func gameLoopFlow() (bool, error) {
	//normal day loop
	dayFlowLoop()
	_, err := strconv.Atoi("-42") //error cheat
	cont := check.GameEndCheck()
	return cont, err
}

func dayFlowLoop() {
	cont := true
	for cont == true {
		models.ShowTime() //test
		cont = getChoices()
		check.StatZeroCheck()
	}

	//after day ends

	endOfWeek() //change into a checker/event
}

func getChoices() bool {
	//check schedule/time events
	ct := models.GetTime()
	cont := true
	//work or free
	switch ct.DayType {
	case 0: //daytime
		fmt.Println("Daytime, oddly nothing to do.")
	case 2: //nightime
		/*if check.UnemployedCheck() == true {
		action.UnemployedActions()}*/
		options, m1 := action.FreeTimeActions() //implement basic menu now
		//display options
		r1 := inputs.StringarrayInput(options)
		cont = action.FreeTimeOutcomes(options[r1-1], m1)
	}
	//stats events
	//options = append...

	//location events
	if check.EventLocationCheck() == true {
		action.EventLocationAction()
	}
	return cont
}

func schedulePlanner() { //not used currently
	fmt.Println("What would you like to do this week?")
	//daytimeoptions := []string{"Go to work", "Stay home", "Go on holiday"}
	//weeknightoptions := []string{"Stay home", "Go to gym"}
}

//move to checker/events
func endOfWeek() {
	//ct := models.GetTime()
	/*if cd[1] == 5 && cw[1] == 1 {
		fmt.Println("How are you feeling about your life choices?")
		//resets schedule to Job schedule. cheat, need an exception for Weekday morning?
		models.SetSchedule(models.GetCurrentPlayer().CurJorb.Schedule)
	}*/
	//look at traits/stats?
}

func endFlow() {
	fmt.Println("It's all over.")
}
