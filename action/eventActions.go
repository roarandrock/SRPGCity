package action

import (
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"fmt"
)

/*
Events to make:
Story events that trigger when goals are achieved, i.e. hitting stat targets
*/

func EventLocationAction() {
	eL1 := models.EventGetByLoc(models.GetPlayerLoc())
	e1 := eL1[0] //cheating
	rn1 := Brng(100)
	if e1.Chance >= rn1 {
		eventActions(e1)
	} else {
		fmt.Println("Test. No event:", rn1)
	}
}

func eventActions(e1 models.EventRandom) {
	switch e1.Nummer {
	case 1:
		fmt.Println("Welcome to East. It's a pricey part of the city. Was working class a few decades ago. " +
			"Good place for restaurants.")
	}
	e1.Used = true
	models.EventUpdate(e1)
}

func UnemployedActions() {
	fmt.Println("You don't have a job. What's the plan?")
	op1 := []string{"Hunt", "Scavenge", "Panic", "Relax"}
	r1 := inputs.StringarrayInput(op1)
	switch r1 {
	case 1:
		n1, jList := models.JorbList()
		fmt.Println("There are:", n1, "jobs currently available in the city.")
		op2 := []string{}
		for _, v := range jList {
			skip := false
			for _, u := range op2 {
				if v.Industry == u {
					skip = true
					break
				}
			}
			if !skip {
				op2 = append(op2, v.Industry)
			}
		}
		fmt.Println("Which industry are you interested in?")
		r2 := inputs.StringarrayInput(op2)
		jList2 := []models.Jorb{}
		op3 := []string{}
		for n2, v := range jList {
			if op2[r2-1] == v.Industry {
				jList2 = append(jList2, v)
				op3 = append(op3, v.Title)
				fmt.Println("Job Option", n2+1)
				models.JorbDisplay(v)
			}
		}
		fmt.Println("Which job would you like to apply for?")
		r3 := inputs.StringarrayInput(op3)
		j1 := jList2[r3-1] //chosen job
		cont1 := models.MinStatCheck(j1.Minstats.MinStats)
		if cont1 == false {
			fmt.Println("Test, stat fail", j1.Minstats.MinStats, models.GetPlayerStats())
		}
		cont2 := models.MinTraitCheck(j1.Minstats.MinTraits)
		if cont2 == false {
			fmt.Println("Test, trait fail", j1.Minstats.MinTraits)
		}
		if cont1 == true && cont2 == true {
			fmt.Println("You got the job!")
			np := models.GetCurrentPlayer()
			np.CurJorb = j1
			models.UpdateCurrentPlayer(np) //updates player with new job
			models.PlayerJorbUpdate()      //sets schedule and traits
		} else {
			fmt.Println("You didn't get the job.")
		}
		//cheat how to show how time consuming this is?
		//takes Energy and hygiene too
		huntImp := models.PlayerMeters{Energy: -10, Stress: 15}
		models.MeterChange(huntImp)
	case 2:
		//dumpster diving, wandering. use this elsewhere?
		fmt.Println("You wander the city looking. What are you looking for?")
		op2 := []string{"Food", "Fun", "Flesh"}
		r2 := inputs.StringarrayInput(op2)
		switch r2 {
		case 1:
			fmt.Println("You find a bagel in a dumpster.")
			r3 := inputs.StringarrayInput([]string{"Eat it", "Gross"})
			switch r3 {
			case 1:
				fmt.Println("Tasty and free.")
				npt := models.PlayerTraits{Lucky: 70}
				npm := models.PlayerMeters{Energy: 15, Stress: -5, Hygiene: -5}
				nps := models.PlayerStats{Knowledge: 5}
				models.TraitChange(npt)
				models.MeterChange(npm)
				models.StatChange(nps)
			case 2:
				fmt.Println("It goes back to going to waste")
				npm := models.PlayerMeters{Energy: -10}
				models.MeterChange(npm)
			}
		case 2:
			fmt.Println("You search for fun.")
			if Brng(100) >= 70 {
				fmt.Println("You find a little park and some company is doing an ice cream promotion.")
				npt := models.PlayerTraits{Lucky: 70}
				npm := models.PlayerMeters{Energy: 10, Stress: -10}
				nps := models.PlayerStats{Knowledge: 5}
				models.TraitChange(npt)
				models.MeterChange(npm)
				models.StatChange(nps)
			} else {
				fmt.Println("You find no fun.")
				npm := models.PlayerMeters{Energy: -5, Stress: -5}
				models.MeterChange(npm)
			}
		case 3: //flesh
			fmt.Println("You talk to strangers until you meet someone interesting.")
			if Brng(100) >= 60 {
				fmt.Println("They invite you back to their place.")
				npt := models.PlayerTraits{Lucky: 60, Smooth: 80}
				npm := models.PlayerMeters{Stress: -20, Soul: 1}
				nps := models.PlayerStats{Connections: 5, Experience: 5}
				models.TraitChange(npt)
				models.MeterChange(npm)
				models.StatChange(nps)
			} else {
				fmt.Println("They get distracted and you never see them again.")
				npm := models.PlayerMeters{Energy: -5}
				models.MeterChange(npm)
			}
		}
	case 3:
		//freak out, lower stats
		fmt.Println("You tell yourself:" +
			"\n\"What are you doing with your life? You'll never find a job! You'll die in the streets!\"")
		npm := models.PlayerMeters{Stress: 20, Soul: -1}
		models.MeterChange(npm)
	case 4:
		//chill
		fmt.Println("You tell yourself:" +
			"\n\"It is all going to be fine. This a chance to find something you love. To pursue your passions.\"")
		npm := models.PlayerMeters{Stress: -5, Soul: 1}
		models.MeterChange(npm)
	}
}
