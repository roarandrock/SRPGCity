package models

/*
Every district has slots for food:
Fast:
Healthy:
Fancy:
Hip:

One resto for each slot. Can play with this, one place has Fast: Taco Hut, another place has Hip: Taqeuira Fresca
*/

type District struct {
	Nummer int
	Name   string
	Place1 FunPlace
}

//move to funplaces
type LocationAction struct {
	Name   string
	Chance int
}

type FunPlace struct {
	Nummer    int
	Name      string
	MinTraits PlayerTraits
	Action1   []LocationAction
}

var cityMap = map[string]District{}

func SetCity() {
	//places
	taco := FunPlace{1, "Taco Feast", PlayerTraits{Income: 10}}
	pizza := FunPlace{2, "Pizzaholic", PlayerTraits{Income: 30}}
	//districts
	West := District{1, "West", taco}
	East := District{2, "East", pizza}
	CityUpdate(West)
	CityUpdate(East)
}

func CityUpdate(cd District) {
	cm := cityMap
	cm[cd.Name] = cd
}

//DistrictGetByName grabs current City by short name
func DistrictGetByName(c string) District {
	cm := cityMap
	i := cm[c]
	return i
}

func FunPlaceGetByName(cd string, FunP string) FunPlace {
	cm := cityMap
	d := cm[cd]
	cFunP := FunPlace{Name: "Nowhere"}
	if d.Place1.Name == FunP {
		cFunP = d.Place1
	}
	return cFunP
}

//DistrictGetByLoc grabs District by location
func DistrictGetByLoc(l int) District {
	cm := cityMap
	cd := District{}
	for _, v := range cm {
		if v.Nummer == l {
			cd = v
		}
	}
	return cd
}

func TravelOptions(l int) []District {
	td := []District{}
	if l == 1 { //cheating
		nd := DistrictGetByLoc(2)
		td = append(td, nd)
	} else if l == 2 {
		nd := DistrictGetByLoc(1)
		td = append(td, nd)
	}
	return td
}
