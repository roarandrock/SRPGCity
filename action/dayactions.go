package action

import (
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"fmt"
)

//move to own file, split name and impact and use id to reference(like mini DB)
type freeTimeAct struct {
	actionname string
	desc       string //description
	deltatime  int
	loc        int //allowed location
	minstats   models.PlayerStats
	mintraits  models.PlayerTraits
	minmeters  models.PlayerMeters
	impstats   models.PlayerStats
	imptraits  models.PlayerTraits
	impmeters  models.PlayerMeters
}

var (
	nothing = freeTimeAct{"Nothing",
		"For when there's nothing better to do",
		1,
		0,
		models.PlayerStats{},
		models.PlayerTraits{},
		models.PlayerMeters{},
		models.PlayerStats{},
		models.PlayerTraits{},
		models.PlayerMeters{}}
	watchTV = freeTimeAct{"Watch TV",
		"Watching the latest in mass communication.",
		1,
		0, //home?
		models.PlayerStats{Stuff: 20}, //mins
		models.PlayerTraits{Independence: 10},
		models.PlayerMeters{Energy: 10},
		models.PlayerStats{Brands: 5}, //impact
		models.PlayerTraits{},
		models.PlayerMeters{Energy: 5, Stress: -20, Hygiene: -20}} //for testing
	sleep = freeTimeAct{"Sleep",
		"ZZZZZZZZZZZZ",
		0,
		0,
		models.PlayerStats{},                                  //mins
		models.PlayerTraits{Alcoholism: 90, Independence: 10}, //need a drink to sleep
		models.PlayerMeters{},
		models.PlayerStats{}, //impact
		models.PlayerTraits{Beauty: 5, Smooth: 5},
		models.PlayerMeters{Energy: 40, Stress: -20, Hygiene: 5, Soul: 1}}
	party = freeTimeAct{"Party", //too vague, and not something that can be done alone. Only good for testing
		"Party!",
		2,
		0,
		models.PlayerStats{}, //mins
		models.PlayerTraits{Independence: 10},
		models.PlayerMeters{Energy: 20, Stress: 80},
		models.PlayerStats{Connections: 5}, //impact
		models.PlayerTraits{Alcoholism: 10, Smooth: 5},
		models.PlayerMeters{Energy: -30, Stress: -10, Hygiene: -25, Soul: 1}}
	shower = freeTimeAct{"Shower",
		"Scrub scrub scrub",
		1,
		0,
		models.PlayerStats{Stuff: 5}, //mins
		models.PlayerTraits{},
		models.PlayerMeters{},
		models.PlayerStats{}, //impact
		models.PlayerTraits{},
		models.PlayerMeters{Energy: -5, Stress: -15, Hygiene: 40}}
)

//FreeTimeActions determines actions in freetime
func FreeTimeActions() ([]string, int) {
	//need to implement general menu options for free time
	m1 := freeMenu()
	//this gets the list of individual actions
	options := []string{}
	switch m1 {
	case 1: //meet people
		fpList := funPlaceList()
		for _, v := range fpList {
			options = append(options, v.Name)
		}
		fmt.Println("Where to meet people?")
	case 2: //Maintenance
		ftlist := freeActList()
		for _, v := range ftlist {
			options = append(options, v.actionname)
		}
	case 3: //Mobility
		dList := models.TravelOptions(models.GetPlayerLoc())
		for _, v := range dList {
			options = append(options, v.Name)
		}
		fmt.Println("You are in the ", models.DistrictGetByLoc(models.GetPlayerLoc()).Name, "district.")
		fmt.Println("Where would you like to go?")
	case 4: //Manipulation
		options = []string{"Quit Game", "Quit this menu"}
		/*
			if models.GetCurJorb().Title != "Unemployed" {
				options = append(options, "Quit Job")
			} */
	}

	return options, m1
}

func freeMenu() int {
	options := []string{"Meet", "Maintenance", "Mobility", "Manipulation"} //add manipulation option, also need to start storing options eslewhere
	r1 := inputs.StringarrayInput(options)
	return r1
}

//FreeTimeOutcomes takes in a string and shows the outcome
func FreeTimeOutcomes(choption string, m1 int) bool {
	cont := true //for day loop
	switch m1 {
	case 1:
		cd := models.DistrictGetByLoc(models.GetPlayerLoc())
		//get place bsed on choption
		fp1 := models.FunPlaceGetByName(cd.Name, choption)
		if models.MinTraitCheck(fp1.MinTraits) == true {
			fmt.Println("You go to", fp1.Name, "to find people.")
		} else {
			fmt.Println("Test: You cannot afford", fp1.Name)
		}
		//cheat for time it takes
		dt := 1
		cont = models.UpdateTimeHour(dt)
		if cont == false {
			fmt.Println("You stayed up the whole night without sleeping")
			sleepDeprived(0)
		}
	case 2:
		ftlist := freeActList()
		ftchosen := freeTimeAct{}
		for _, v := range ftlist {
			if v.actionname == choption {
				ftchosen = v
			}
		}
		//fmt.Println("Test 1:", ftchosen.imptraits)
		models.StatChange(ftchosen.impstats)
		models.TraitChange(ftchosen.imptraits)
		models.MeterChange(ftchosen.impmeters)
		/*if ftchosen.deltatime > 1 {
			models.GetenSetSchedEvent("Sleeping") //generic placeholder
			cd, cw := models.GetFutureTime(1)
			nSch := 0 //setting next segment to busy
			models.UpdateSchedPerTime(cd, cw, nSch)
		}*/
		//need to deplete the time meter
		if ftchosen.deltatime > 0 {
			cont = models.UpdateTimeHour(ftchosen.deltatime)
			if cont == false {
				fmt.Println("You stayed up the whole night without sleeping")
				sleepDeprived(0)
			}
		}
		//makes changes if sleep is chosen
		if ftchosen.actionname == "Sleep" {
			cont = false
			sH := models.SleepTime()
			sleepDeprived(sH)
		}
	case 3:
		nd := models.DistrictGetByName(choption)
		cp := models.GetCurrentPlayer()
		cp.Loc = nd.Nummer
		fmt.Println("You go to:", nd.Name)
		models.UpdateCurrentPlayer(cp)
		//need to show time this takes, assume 1 hour for now. need to change int to float for decimals
		dt := 1
		cont = models.UpdateTimeHour(dt)
		if cont == false {
			fmt.Println("You stayed up the whole night without sleeping")
			sleepDeprived(0)
		}
	case 4:
		switch choption {
		case "Quit Job":
			fmt.Println("You quit your job.")
			np := models.GetCurrentPlayer()
			sj := models.Jorb{Title: "Unemployed", CompanyName: "Nowhere"}
			np.CurJorb = models.JorbGet(sj)
			models.UpdateCurrentPlayer(np)
			models.PlayerJorbUpdate()
		case "Quit Game":
			ct := models.GetTime() //cheat
			ct.DayN = 100
			models.UpdateTime(ct)
			cont = false
		case "Quit this menu":
			fmt.Println("Ok")
		}
	}
	return cont
}

//FreeActList makes a list of valid actions
func freeActList() []freeTimeAct {
	//get all possible actions
	ftlist1 := []freeTimeAct{watchTV, sleep, party, shower} //needs a function
	//need a default action for when nothing else works
	ftlist2 := []freeTimeAct{}
	//check time, location, stats
	for _, v := range ftlist1 {
		if models.MinMeterCheck(v.minmeters) == true {
			ftlist2 = append(ftlist2, v)
		} else {
			fmt.Println("Test: You cannot", v.actionname)
		} //can do the same for traits and stats if this works
	}
	return ftlist2
}

func funPlaceList() []models.FunPlace {
	fpList1 := []models.FunPlace{}
	cd := models.DistrictGetByLoc(models.GetPlayerLoc())
	fpList1 = append(fpList1, cd.Place1) //cheating
	return fpList1
}

func sleepDeprived(sH int) {
	fmt.Println("You got", sH, "hours of sleep.")
	//need to implement impact
}

/*
//WorkOutcomes is the impact from chosing the action
func WorkOutcomes(choption string) {
	setWorkAct() //move?
	wtlist := []workTimeAct{work, sleep1}
	wtchosen := workTimeAct{}
	for _, v := range wtlist {
		if v.actionname == choption {
			wtchosen = v
		}
	}
	models.MeterChange(wtchosen.deltameters)
	chance1 := Brng(100) //different function
	if chance1 <= wtchosen.firedchance {
		fmt.Println("You're fired.")
		np := models.GetCurrentPlayer()
		sj := models.Jorb{Title: "Unemployed", CompanyName: "Nowhere"}
		np.CurJorb = models.JorbGet(sj)
		models.UpdateCurrentPlayer(np)
		models.PlayerJorbUpdate()
	}
}

type workTimeAct struct {
	actionname  string
	firedchance int
	deltameters models.PlayerMeters //get from jorb definition, except for special actions
	//promotion chances, getting fired chances
}

var (
	work = workTimeAct{"Working for the machine.",
		10,
		models.PlayerMeters{}} //doesn't pull them in at start
	sleep1 = workTimeAct{"Getting paid to nap. The perfect crime.",
		70,
		models.PlayerMeters{Energy: 25}}
)

//need this
func setWorkAct() {
	work.deltameters = models.GetCurJorb().Impact.ImpactMeters
}

//WorkActions chooses actions at work
func WorkActions() []string {
	cj := models.GetCurJorb()
	fmt.Println("You go to work at ", cj.CompanyName)
	wtlist := []workTimeAct{work, sleep1} //need a function for this
	options := []string{}
	for _, v := range wtlist {
		options = append(options, v.actionname)
	}
	return options
}

*/
