package narrative

import (
	"UrbanRPG1/inputs"
	"UrbanRPG1/model"
	"fmt"
)

//Intro sets up the game, from a story perspective
func Intro() {
	fmt.Println("First day in the big city." +
		"\nYou're on the sidewalk. Outside your arguably affordable apartment." +
		"\nA man talks to you. He's wearing small, metallic glasses and a heavy black coat.")
	fmt.Println("\"Who are you? What do you want? What will you be?" +
		"\nThis place is a machine.\"")
	options := []string{"A cog", "A wrench", "A hammer"}
	r1 := inputs.StringarrayInput(options)
	switch r1 {
	case 1:
		fmt.Println("A vital piece. Fitting into the established motions. Small and replacable.")
		models.SetGoals(1)
	case 2:
		fmt.Println("Fixing, adjusting, changing. You are no cog. Still you serve the machine.")
		models.SetGoals(2)
	case 3:
		fmt.Println("Smashing, breaking. Violent and destructive. A threat to order, efficiency and the machine itself.")
		models.SetGoals(3)
	}
	fmt.Println("New job starts tomorrow.")
}
