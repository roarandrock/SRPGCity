package check

import (
	"UrbanRPG1/model"
	"fmt"
)

//ErrorCheck for checking errors
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}
}

//GameEndCheck checks for end conditions
func GameEndCheck() bool {
	cont := goalAchievedCheck()
	ct := models.GetTime()
	if ct.DayN >= 10 {
		cont = false
	}
	return cont
}

func goalAchievedCheck() bool {
	cont := true
	pg := models.GetGoals()
	pstats := models.GetPlayerStats()
	ptraits := models.GetPlayerTraits()
	if pstats.Connections == pg.MinStats.Connections { //make a compare stats function
		cont = false //need victory text written for each goal
	}
	if ptraits.Income == pg.MinTraits.Income {
		cont = false
	}
	if ptraits.Independence == pg.MinTraits.Independence {
		cont = false
	}
	if cont == false {
		fmt.Println(pg.Victory)
	}
	return cont
}

//StatZeroCheck checks if a vital stat is zero
func StatZeroCheck() bool {
	cont := true
	cpt := models.GetPlayerTraits()
	cpm := models.GetPlayerMeters()
	//fmt.Println("test, indp", cpt.Independence, cpt)
	if cpm.Energy <= 0 {
		fmt.Println("You have no energy.")
		fmt.Println("You pass out, unconcious.")
		/*
			cd, cw := models.GetFutureTime(1)
			models.UpdateSchedPerTime(cd, cw, 0)
			models.GetenSetSchedEvent("You went too far, too hard. Need rest.")
		*/
		cpm.Energy = 50
	}
	if cpm.Stress >= 100 {
		fmt.Println("You are so stressed you suffer a breakdown.")
		/*
			cd, cw := models.GetFutureTime(1)
			models.UpdateSchedPerTime(cd, cw, 0)
			models.GetenSetSchedEvent("You need time to recover.")
		*/
		cpm.Stress = 1
	}
	if cpm.Hygiene <= 0 {
		fmt.Println("You are a filthy animal. You are quarantined in your apartment.")
		/*
			cd, cw := models.GetFutureTime(1)
			models.UpdateSchedPerTime(cd, cw, 0)
			models.GetenSetSchedEvent("You need extra time to clean yourself. So dirty.")
		*/
		cpm.Hygiene = 5
	}
	if cpt.Independence <= 0 {
		fmt.Println("You have no independence. Game over.")
		cont = false
	}
	models.UpdatePlayerMeters(cpm)
	return cont
}

func EventLocationCheck() bool {
	cont := false
	eventList := models.EventGetByLoc(models.GetPlayerLoc())
	if len(eventList) != 0 {
		cont = true
		fmt.Println("Test, there are events here:", eventList)
	}
	return cont
}

func UnemployedCheck() bool {
	check1 := false
	cp := models.GetCurrentPlayer()
	if cp.CurJorb.Title == "Unemployed" {
		check1 = true
	}
	return check1
}
