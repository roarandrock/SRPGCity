package models

//Goal defines player goals
type Goal struct {
	Descr     string
	MinStats  PlayerStats
	MinTraits PlayerTraits
	Victory   string
}

var (
	mygoal = Goal{}
	goal1  = Goal{"Succeed at fitting into the city",
		PlayerStats{Connections: 100},
		PlayerTraits{},
		"You have many friends, acquaintances, and connections."}
	goal2 = Goal{"Rule over others in the city",
		PlayerStats{},
		PlayerTraits{Income: 100},
		"You have become wealthy. And this city, this is the easiest form of power."}
	goal3 = Goal{"Break the city before it breaks you",
		PlayerStats{},
		PlayerTraits{Independence: 100},
		"You are free."}
)

//SetGoals takes in an integer to set the goal
func SetGoals(n int) {
	gray := []Goal{goal1, goal2, goal3}
	mygoal = gray[n-1]
}

func GetGoals() Goal {
	exportg := mygoal
	return exportg
}
