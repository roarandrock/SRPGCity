package models

/*

Location, can be specific district, or anywhere (0)

Specific actions/effects, can occur within the event?

Maybe classify by kind of event? Stat, location, story...

*/

type EventRandom struct {
	Nummer int
	Name   string
	Loc    int
	Chance int  //change to trigger event
	Used   bool //change after even is triggered
}

var eventRMap = map[string]EventRandom{}

func SetEvents() {
	//set Events:
	event1 := EventRandom{1, "First time in Oost", 2, 100, false}
	//Store them:
	EventUpdate(event1)
}

func EventUpdate(cd EventRandom) {
	cm := eventRMap
	cm[cd.Name] = cd
}

func EventGetByName(c string) EventRandom {
	cm := eventRMap
	i := cm[c]
	return i
}

//none of these right now
func EventGetByLoc(l int) []EventRandom {
	cm := eventRMap
	cd := []EventRandom{}
	for _, v := range cm {
		if v.Loc == l && v.Used == false {
			cd = append(cd, v)
		}
	}
	return cd
}
