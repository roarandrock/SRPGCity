package models

import "fmt"

/*
to keep and modify schedule

Weekday, weekend
Morning, Daytime, Evening, Night, Midnight, Dawn

Both split into arrays
Week = 0 weekday, 1 weekend
Day = 0 Morning- 5 dawn
Weekday, weekend
Morning, Daytime, Evening, Night, Midnight, Dawn



freetime = 1
worktime = 2
busy = 0

40 hour work week:
0,2,1,1,1,1,1,1,1,1,1,1

*/

var (
	wd0 = 1 //weekday morning
	wd1 = 1
	wd2 = 1
	wd3 = 1
	wd4 = 1
	wd5 = 1
	we0 = 1
	we1 = 1
	we2 = 1
	we3 = 1
	we4 = 1
	we5 = 1
)

var (
	sevent0 = "Getting ready for work."
)

//SetSchedule takes in new schedule
func SetSchedule(newsc []int) {
	//can make a check then return error
	if len(newsc) != 12 {
		fmt.Println("Schedule error, wrong size", len(newsc))
		return
	}
	wd0 = newsc[0]
	wd1 = newsc[1]
	wd2 = newsc[2]
	wd3 = newsc[3]
	wd4 = newsc[4]
	wd5 = newsc[5]
	we0 = newsc[6]
	we1 = newsc[7]
	we2 = newsc[8]
	we3 = newsc[9]
	we4 = newsc[10]
	we5 = newsc[11]
}

//GetSched returns current schedule
func GetSched() []int {
	exports := []int{wd0, wd1, wd2, wd3, wd4, wd5,
		we0, we1, we2, we3, we4, we5}
	return exports
}

func GetenSetSchedEvent(newM string) string {
	export := sevent0
	if newM != "" {
		sevent0 = newM
	}
	return export
}

func UpdateSchedPerTime(cd []int, cw []int, nSch int) {
	oldSch := GetSched() //do we need a temp schedule for busy events?
	//fmt.Println("Test old:", oldSch)
	if cw[1] == 0 { //weekday
		oldSch[cd[1]] = nSch
	} else if cw[1] == 1 { //weekend
		oldSch[cd[1]+6] = nSch
	}
	SetSchedule(oldSch)
	//fmt.Println("Test new:", GetSched())
}
