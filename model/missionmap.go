package models

import "fmt"

/*
Try and draw the place

Think of Descent- lots of little pieces coming together to make a map

Keep it simple: All you need is the tile for the player, and the next tile

No big grid, just those two. The coordinates within the tiles are where things are placed

Should focus on coordinates within a room. Characters track the Room ID and their coordiantes within it

*/

type Room struct {
	Id     int
	Width  int
	Height int
	//doors
}

type Tile struct {
	x int
	y int
	f int //there are four sides, which sides need shading
	/*
		   0 is blank
			 1 is wall /
			 2 is door
			 8 is NPC
			 9 is C (need multiple for team...)
	*/
}

var (
	Room1  = Room{1, 8, 3}
	Room2  = Room{2, 10, 10}
	tStack = []Tile{}
)

//MakeMap is the main flow for making the mission map
func MakeMap() {
	//my table is one blank space
	//pick rooms
	pick1 := Room1
	pick2 := Room2
	//make empty map
	blankMmap(pick1, pick2)
	//draw rooms
	coord := []int{0, 0}
	coord = drawRoom(pick1, coord)
	drawRoom(pick2, coord)
	doorSet(pick1, pick2)
	//fmt.Println("test:", tStack) //test
	//drawRoom(pick2, coord, MPrint)
	//draw final map, should not be in models but in viewer
	//can place dummy dudes inside for testing
	drawMMap(pick1, pick2)
}

//cheatin
func doorSet(r1 Room, r2 Room) {
	nt, _, _ := tileFind(r1.Width-1, 1)
	nt.f = 2
	tileSet(nt.x, nt.y, nt)
	nt, _, _ = tileFind(r1.Width, 1)
	nt.f = 2
	tileSet(nt.x, nt.y, nt)
}

//make a blank mission map using the room sizes
func blankMmap(pick1 Room, pick2 Room) {
	xM := pick1.Width + pick2.Width
	yM := pick1.Height + pick2.Height
	xl := 0
	yl := 0
	tStack = []Tile{}
	for i := 0; i < yM; i++ { //makes all the blank tiles I need
		for i2 := 0; i2 < xM; i2++ {
			bTile := Tile{xl, yl, 0}
			tStack = append(tStack, bTile)
			xl++
		}
		xl = 0
		yl++
	}
	//fmt.Println("test:", tStack) //test
}

func tileFind(x0 int, y0 int) (Tile, int, bool) {
	bTile := Tile{}
	tl := 0
	tcheck := false
	for i, v := range tStack {
		if v.x == x0 {
			if v.y == y0 {
				bTile = v
				tl = i
				tcheck = true
			}
		} //add else err check
	}
	//search tStack for Tile based on coords
	return bTile, tl, tcheck
}

func tileSet(x0 int, y0 int, nt Tile) bool {
	_, oi, tcheck := tileFind(x0, y0)
	tStack[oi] = nt
	goed := tcheck
	return goed
}

func drawRoom(r1 Room, coord []int) []int {
	x0 := coord[0] //starting coordinates, actual coordinates
	y0 := coord[1]
	xl := 0 //local x,y
	yl := 0
	h := r1.Height
	w := r1.Width
	//update tiles based on rooms
	//set from coordinate 0, all the way right, then reset, up, then reset
	for i2 := 0; i2 < h; i2++ {
		for i3 := 0; i3 < w; i3++ {
			nt, _, _ := tileFind(x0, y0)
			if i2 == 0 || i2 == h-1 {
				nt.f = 1
			} else {
				if i3 == 0 || i3 == w-1 {
					nt.f = 1
				} else {
					nt.f = 0
				}
			} //checks if on starting row or top row
			tileSet(x0, y0, nt)
			//fmt.Println("tile test", nt)
			x0++
			xl++
		}
		x0 = coord[0]
		xl = 0
		y0++
		yl++
	}
	//add NPCs, cheating

	cMon := GetMonById(r1.Id)
	mPos := cMon.Coord
	offx := mPos[0] + coord[0]
	offy := mPos[1] + coord[1]
	mt, _, _ := tileFind(offx, offy) //offset, can make a function
	mt.f = 8
	tileSet(mt.x, mt.y, mt)
	//add PC
	cp := GetCurrentPlayer()
	if cp.RoomId == r1.Id {
		pPos := cp.Coord
		offx = pPos[0] + coord[0]
		offy = pPos[1] + coord[1]
		pt, _, _ := tileFind(offx, offy) //offset, can make a function
		pt.f = 9
		tileSet(pt.x, pt.y, pt)
	}
	//coord defines starting and stopping point of the room. how to choose it for the next one?
	//cheating!
	newcoord := []int{w, 0}
	return newcoord
}

//drawMMap should be moved, for drawing the map
func drawMMap(pick1 Room, pick2 Room) {
	x0 := 0
	y0 := 0 //need to start on top
	if pick1.Height > pick2.Height {
		y0 = pick1.Height - 1
	} else {
		y0 = pick2.Height - 1
	}
	xMax := pick1.Width + pick2.Width
	yMax := y0 + 2 //not sure why
	//options
	empty := 0
	wall := 1
	door := 2
	npc := 8
	pC := 9
	//fuck it, everything should be /s or |, same symbol for all walls
	//fmt.Println("test:", tStack) //test
	for i3 := 0; i3 < yMax; i3++ {
		for i2 := 0; i2 < xMax; i2++ { //make top row down, printing, wrap to move down
			ct, _, _ := tileFind(x0, y0)
			//fmt.Print(ct) // test
			switch ct.f { //make a separate function?
			case wall:
				fmt.Print("-")
			case empty:
				fmt.Print(" ")
			case door:
				fmt.Print("|")
			case npc:
				//need a get monster? table!
				fmt.Print("M")
			case pC:
				fmt.Print("P")
			}
			x0++
		}
		x0 = 0
		fmt.Print("\n")
		y0--
	}
	//MPrint := []string{}
	/*
		for i := 0; i < h0; i++ {
			fmt.Print(MPrint[i], "\n")
		}
	*/
}

//move to checkers?
//What about the second room? Player uses local coordinates. Should verify room
func CollisionCheck(cCoord []int, roomId int) []int {
	freeC := []int{}
	//N - 1, S - 2, W - 3, E - 4
	x0 := cCoord[0]
	y0 := cCoord[1]
	if roomId == 2 {
		fmt.Println("Collision may break, in room 2")
		x0 = x0 + Room1.Width
		y0 = y0 + 0
	}
	t1, _, tCheck := tileFind(x0, y0+1) //checking North
	if tCheck == true {
		if t1.f == 0 {
			freeC = append(freeC, 1)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	t1, _, tCheck = tileFind(x0, y0-1) //checking south
	if tCheck == true {
		if t1.f == 0 {
			freeC = append(freeC, 2)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	t1, _, tCheck = tileFind(x0-1, y0) //checking west
	if tCheck == true {
		if t1.f == 0 {
			freeC = append(freeC, 3)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	t1, _, tCheck = tileFind(x0+1, y0) //checking east
	//fmt.Println("east", t1, tCheck)
	if tCheck == true {
		if t1.f == 0 {
			freeC = append(freeC, 4)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	//fmt.Println(freeC)
	return freeC
}

//move to checkers?
//how to tell which monster is which? Need a find monster by location?
//cheating for range, assumes exact range
func RangeCheck(cCoord []int, roomId int, attkRange int) []Monster {
	monIds := []Monster{}
	//N - 1, S - 2, W - 3, E - 4
	x0 := cCoord[0]
	y0 := cCoord[1]
	if roomId == 2 {
		fmt.Println("Range may break, in room 2")
		x0 = x0 + Room1.Width
		y0 = y0 + 0
	}
	t1, _, tCheck := tileFind(x0, y0+attkRange) //checking North
	if tCheck == true {
		if t1.f == 8 {
			lCoord := []int{x0, y0 + attkRange}
			lMon := GetMonByLoc(lCoord, roomId)
			monIds = append(monIds, lMon)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	t1, _, tCheck = tileFind(x0, y0-attkRange) //checking south
	if tCheck == true {
		if t1.f == 8 {
			lCoord := []int{x0, y0 - attkRange}
			lMon := GetMonByLoc(lCoord, roomId)
			monIds = append(monIds, lMon)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	t1, _, tCheck = tileFind(x0-attkRange, y0) //checking west
	if tCheck == true {
		if t1.f == 8 {
			lCoord := []int{x0 - attkRange, y0}
			lMon := GetMonByLoc(lCoord, roomId)
			monIds = append(monIds, lMon)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	t1, _, tCheck = tileFind(x0+attkRange, y0) //checking east
	//fmt.Println("east", t1, tCheck)
	if tCheck == true {
		if t1.f == 8 {
			lCoord := []int{x0 + attkRange, y0}
			lMon := GetMonByLoc(lCoord, roomId)
			monIds = append(monIds, lMon)
		} else if t1.f == 2 {
			fmt.Println("Door, do something cheat")
		}
	}
	//fmt.Println(monIds)
	//use get monster by location to get actual monsterids and return those

	return monIds
}
